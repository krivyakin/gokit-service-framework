package middleware

import (
	"github.com/google/uuid"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"net/http"
)

type HTTPMiddleware func(next http.Handler) http.Handler

type loggingMiddleware struct {
	logger log.Logger
	next   http.Handler
}

func NewLoggingMiddleware(logger log.Logger) HTTPMiddleware {
	logger = logger.WithLocation("http.loggingMiddleware")
	return func(next http.Handler) http.Handler {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (l *loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var lrw *loggingResponseWriter
	var status int
	defer func() {
		l.logger.Info("status", status, "request", r.RequestURI)
	}()
	ctx := log.ContextWithReqID(r.Context(), uuid.New().String())
	r = r.WithContext(ctx)

	lrw = NewLoggingResponseWriter(w)
	l.next.ServeHTTP(lrw, r)
	status = lrw.statusCode
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
