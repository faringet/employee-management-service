package transport

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"

	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
)

func MakePingHandler(logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	return kithttp.NewServer(
		endpoints.MakePingEndpoint(),
		pingDecodeRequest,
		encodeResponse,
		opts...,
	)
}

func pingDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
