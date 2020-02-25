package http

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"github.com/krivyakin/gokit-service-framework/pkg/service1/transport"
	"net/http"
)

func NewServiceRouter(_ context.Context, svcEndpoints transport.Endpoints, logger log.Logger) *mux.Router {
	logger = logger.WithLocation("http.service")
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger.KitLogger()),
		kithttp.ServerErrorEncoder(encodeErrorResponse),
	}

	r.Methods("GET").Path("/config").Handler(kithttp.NewServer(
		svcEndpoints.Config,
		decodeConfigRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodeConfigRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return transport.ConfigRequest{}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
