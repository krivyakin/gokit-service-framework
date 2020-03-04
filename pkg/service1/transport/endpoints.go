package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
)

/*
NEW_HANDLER_STEP5: create an endpoint for a new function. He you need to convert response from service to
transport layer response (see makeConfigEndpoint/makeInfoEndpoint).
*/
type Endpoints struct {
	Config endpoint.Endpoint
	Info   endpoint.Endpoint
}

func MakeEndpoints(s service1.Service) Endpoints {
	return Endpoints{
		Config: makeConfigEndpoint(s),
		Info:   makeInfoEndpoint(s),
	}
}

func makeConfigEndpoint(s service1.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return ConfigResponse{
			// copy the whole response from the service
			Config: s.Config(ctx),
		}, nil
	}
}

func makeInfoEndpoint(s service1.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		info := s.Info(ctx)
		return InfoResponse{
			// use the service response to create a new format of response
			Uptime: info.Uptime.Seconds(),
		}, nil
	}
}
