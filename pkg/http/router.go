package http

import (
	"github.com/gorilla/mux"
	http_middleware "github.com/krivyakin/gokit-service-framework/pkg/http/middleware"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func NewDefaultRouter(logger log.Logger, timeout time.Duration) *mux.Router {
	router := mux.NewRouter()

	router.Use(http_middleware.NewResponseWriterMiddleware())
	router.Use(http_middleware.NewLoggingMiddleware(logger))
	router.Use(http_middleware.NewMetricsMiddleware())
	router.Use(http_middleware.NewTimeoutMiddleware(timeout))

	router.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	return router
}
