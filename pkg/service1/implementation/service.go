package implementation

import (
	"context"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
	"github.com/spf13/viper"
)

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) service1.Service {
	logger = logger.WithLocation("service")
	return &service{
		logger: logger,
	}
}

func (s *service) Config(_ context.Context) map[string]interface{} {
	return viper.AllSettings()
}
