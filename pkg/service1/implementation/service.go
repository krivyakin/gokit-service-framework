package implementation

import (
	"context"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
	"github.com/spf13/viper"
	"time"
)

type service struct {
	logger    log.Logger
	startedAt time.Time
}

func NewService(logger log.Logger) service1.Service {
	logger = logger.WithLocation("service")
	return &service{
		logger:    logger,
		startedAt: time.Now(),
	}
}

// NEW_HANDLER_STEP2: implement a new function from service1.Service interface
func (s *service) Config(_ context.Context) map[string]interface{} {
	return viper.AllSettings()
}

func (s *service) Info(_ context.Context) service1.ServiceInfo {
	return service1.ServiceInfo{
		Uptime: time.Now().Sub(s.startedAt),
	}
}
