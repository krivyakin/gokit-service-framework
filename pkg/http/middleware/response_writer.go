package middleware

import (
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, 0}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type responseWriterMiddleware struct {
	next http.Handler
}

func NewResponseWriterMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return &responseWriterMiddleware{
			next: next,
		}
	}
}

func (l *responseWriterMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var lrw = NewLoggingResponseWriter(w)
	l.next.ServeHTTP(lrw, req)
}
