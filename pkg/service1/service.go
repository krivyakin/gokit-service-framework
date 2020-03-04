package service1

import (
	"context"
	"time"
)

type ServiceInfo struct {
	Uptime time.Duration `json:"uptime"`
}

// NEW_HANDLER_STEP1: add new function to interface
type Service interface {
	Config(ctx context.Context) map[string]interface{}
	Info(ctx context.Context) ServiceInfo
}
