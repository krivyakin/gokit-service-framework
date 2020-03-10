package http

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
	"github.com/krivyakin/gokit-service-framework/pkg/service1/transport"
	"net/http"
)

func RegisterService(service service1.Service, logger log.Logger, r *mux.Router) {
	logger = logger.WithLocation("http.service")
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger.KitLogger()),
		kithttp.ServerErrorEncoder(encodeErrorResponse),
	}

	endpoints := transport.MakeEndpoints(service)
	//NEW_HANDLER_STEP6: add an HTTP handler for a new endpoint
	r.Methods("GET").Path("/config").Handler(kithttp.NewServer(
		endpoints.Config,
		decodeConfigRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/info").Handler(kithttp.NewServer(
		endpoints.Info,
		decodeConfigRequest,
		encodeResponse,
		options...,
	))
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
