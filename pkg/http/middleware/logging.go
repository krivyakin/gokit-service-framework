package middleware

import (
	"github.com/google/uuid"
	http2 "github.com/krivyakin/gokit-service-framework/pkg/http"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"net/http"
)

type loggingMiddleware struct {
	logger log.Logger
	next   http.Handler
}

func NewLoggingMiddleware(logger log.Logger) http2.HTTPMiddleware {
	logger = logger.WithLocation("http.loggingMiddleware")
	return func(next http.Handler) http.Handler {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (l *loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var lrw *http2.LoggingResponseWriter
	var status int
	defer func() {
		l.logger.Info("status", status, "request", r.RequestURI)
	}()
	ctx := log.ContextWithReqID(r.Context(), uuid.New().String())
	r = r.WithContext(ctx)

	lrw = http2.NewLoggingResponseWriter(w)
	l.next.ServeHTTP(lrw, r)
	status = lrw.StatusCode
}