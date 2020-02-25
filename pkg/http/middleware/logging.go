package middleware

import (
	"github.com/google/uuid"
	http2 "github.com/krivyakin/gokit-service-framework/pkg/http"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"net/http"
	"runtime/debug"
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

func (l *loggingMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqid := uuid.New().String()
	ctx := log.ContextWithReqID(req.Context(), reqid)
	req = req.WithContext(ctx)

	var lrw *loggingResponseWriter = w.(*loggingResponseWriter)

	defer func() {
		var keyval []interface{}
		if r := recover(); r != nil {
			lrw.WriteHeader(http.StatusInternalServerError)
			keyval = append(keyval, "error", r, "stack", string(debug.Stack()))
		}
		keyval = append([]interface{}{ "status", lrw.StatusCode, "request", req.RequestURI, "reqid", reqid}, keyval...)

		if lrw.StatusCode == http.StatusOK || lrw.StatusCode == http.StatusNoContent ||
			lrw.StatusCode == http.StatusMovedPermanently || lrw.StatusCode == http.StatusFound {
			l.logger.Infom("REQUEST", keyval...)
		} else {
			l.logger.Errorm("REQUEST", keyval...)
		}
	}()
	l.next.ServeHTTP(lrw, req)
}