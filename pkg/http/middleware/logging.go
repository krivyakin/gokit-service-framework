package middleware

import (
	"github.com/google/uuid"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"net/http"
	"runtime/debug"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next   http.Handler
}

func NewLoggingMiddleware(logger log.Logger) func(next http.Handler) http.Handler {
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

	defer func(start time.Time) {
		var keyval []interface{}
		if r := recover(); r != nil {
			lrw.WriteHeader(http.StatusInternalServerError)
			keyval = append(keyval, "error", r, "stack", string(debug.Stack()))
		}
		basicInfo := []interface{}{
			"status", lrw.StatusCode,
			"request", req.RequestURI,
			"reqid", reqid,
			"elapsed", time.Now().Sub(start).Seconds(),
		}
		keyval = append(basicInfo, keyval...)

		if lrw.StatusCode == http.StatusOK || lrw.StatusCode == http.StatusNoContent ||
			lrw.StatusCode == http.StatusMovedPermanently || lrw.StatusCode == http.StatusFound {
			l.logger.Infom("REQUEST", keyval...)
		} else {
			l.logger.Errorm("REQUEST", keyval...)
		}
	}(time.Now())
	l.next.ServeHTTP(lrw, req)
}