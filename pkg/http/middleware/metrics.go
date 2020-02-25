package middleware

import (
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	http2 "github.com/krivyakin/gokit-service-framework/pkg/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"time"
)

type httpMetrics struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	requestLatencyH metrics.Histogram
}

type metricsMiddleware struct {
	metrics httpMetrics
	next    http.Handler
}

func NewMetricsMiddleware() http2.HTTPMiddleware {
	fieldKeys := []string{"method", "Path", "status"}
	metrics := httpMetrics{
		requestCount: kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "http",
			Subsystem: "service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		requestLatency: kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace:  "http",
			Subsystem:  "service",
			Name:       "request_latency_microseconds",
			Help:       "Total duration of requests in microseconds.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, fieldKeys),
		requestLatencyH: kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: "http",
			Subsystem: "service",
			Name:      "request_latencyH_microseconds",
			Help:      "Total duration of requests in microseconds (hist).",
			Buckets:   []float64{0.5, 0.9, 0.99, 2, 5},
		}, fieldKeys),
	}

	return func(next http.Handler) http.Handler {
		return &metricsMiddleware{
			metrics: metrics,
			next:    next,
		}
	}
}

func (l *metricsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var lrw *loggingResponseWriter = w.(*loggingResponseWriter)

	defer func(begin time.Time) {
		statusStr := strconv.Itoa(lrw.StatusCode)
		lvs := []string{"method", r.Method, "Path", r.URL.Path, "status", statusStr}
		l.metrics.requestCount.With(lvs...).Add(1)
		l.metrics.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		l.metrics.requestLatencyH.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	l.next.ServeHTTP(lrw, r)
}
