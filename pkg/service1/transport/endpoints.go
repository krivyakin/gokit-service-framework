package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
)

type Endpoints struct {
	Config endpoint.Endpoint
}

func MakeEndpoints(s service1.Service) Endpoints {
	return Endpoints{
		Config: makeConfigEndpoint(s),
	}
}

func makeConfigEndpoint(s service1.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return ConfigResponse{
			Config: s.Config(ctx),
		}, nil
	}
}
