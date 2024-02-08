package transport

import (
	"context"
	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"net/http"
	"strconv"

	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func MakeDeleteCommunicationTemplateByIDHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeDeleteCommunicationTemplateByIDEndpoint(s),
		deleteCommunicationTemplateByIDDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func deleteCommunicationTemplateByIDDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		transportLogger.Error("missing status id in path")
		return nil, ie.Error{
			Code:    http.StatusBadRequest,
			Message: "missing status id in path",
		}
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		transportLogger.Error("status id must be an integer")
		return nil, ie.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "request validation failed",
			Details: []ie.Details{
				{
					Field:   "id",
					Message: "status id must be an integer",
				},
			},
		}
	}

	return endpoints.DeleteCommunicationTemplateByIDRequest{ID: id}, nil
}
