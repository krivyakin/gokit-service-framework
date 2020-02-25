package middleware

import "github.com/krivyakin/gokit-service-framework/pkg/service1"

type Middleware func(service service1.Service) service1.Service
