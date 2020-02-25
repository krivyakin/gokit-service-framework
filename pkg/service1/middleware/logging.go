package middleware

import (
	"context"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
)

type loggingMiddleware struct {
	next   service1.Service
	logger log.Logger
}

func NewLoggingMiddleware(logger log.Logger) Middleware {
	logger = logger.WithLocation("service.loggingMiddleware")
	return func(service service1.Service) service1.Service {
		return &loggingMiddleware{
			next:   service,
			logger: logger,
		}
	}
}
func (l *loggingMiddleware) Config(ctx context.Context) map[string]interface{} {
	resp := l.next.Config(ctx)
	l.logger.Infom("config requested", "config", resp)
	return resp
}
