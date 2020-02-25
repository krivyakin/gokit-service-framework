package service1

import "context"

type Service interface {
	Config(ctx context.Context) map[string]interface{}
}
