package transport

import (
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"

	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
)

func MakeEmployeeHandler(logger *zap.Logger, templateService templateservice.Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	endpoint := endpoints.MakeEmployeeEndpoint(templateService)

	return kithttp.NewServer(
		endpoint,
		employeeDecodeRequest,
		encodeResponse,
		opts...,
	)
}
