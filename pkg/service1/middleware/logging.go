package middleware

import (
	"context"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
)

type loggingMiddleware struct {
	middlewareBase
	logger log.Logger
}

func NewLoggingMiddleware(logger log.Logger) Middleware {
	logger = logger.WithLocation("service.loggingMiddleware")
	return func(service service1.Service) service1.Service {
		return &loggingMiddleware{
			middlewareBase: *newBase(service),
			logger:         logger,
		}
	}
}

func (l *loggingMiddleware) Config(ctx context.Context) map[string]interface{} {
	resp := l.middlewareBase.Config(ctx)
	l.logger.Infom("config requested", "config", resp)
	return resp
}

// Uncomment this function if you want to add middleware for Info function
/*
func (l *loggingMiddleware) Info(ctx context.Context) service1.ServiceInfo {
	l.logger.Infom("this is info request")
	return l.middlewareBase.Info(ctx)
}
*/
