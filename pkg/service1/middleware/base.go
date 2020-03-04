package middleware

import (
	"context"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
)

type middlewareBase struct {
	next service1.Service
}

func newBase(next service1.Service) *middlewareBase {
	return &middlewareBase{
		next: next,
	}
}

// NEW_HANDLER_STEP3: add base middleware implementation that just invokes the next middleware
func (e *middlewareBase) Config(ctx context.Context) map[string]interface{} {
	return e.next.Config(ctx)
}

func (e *middlewareBase) Info(ctx context.Context) service1.ServiceInfo {
	return e.next.Info(ctx)
}
