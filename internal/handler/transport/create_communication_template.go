package transport

import (
	"context"
	"encoding/json"
	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func MakeCreateCommunicationTemplateHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeCreateCommunicationTemplateEndpoint(s),
		createCommunicationTemplateDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func createCommunicationTemplateDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateCommunicationTemplateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		transportLogger.Error("decode the request body", zap.Error(err))
		return nil, ie.Error{
			Code:    http.StatusBadRequest,
			Message: "unable to decode the request body",
		}
	}

	requestValidator := validator.New()
	err = requestValidator.Struct(req)
	if err != nil {
		transportLogger.Error("request body is not valid", zap.Error(err))
		errToReturn := ie.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "incorrect request format",
		}
		for _, err := range err.(validator.ValidationErrors) {
			details := ie.Details{
				Field:   err.Field(),
				Message: err.Error(),
			}
			errToReturn.Details = append(errToReturn.Details, details)
		}
		return nil, errToReturn
	}

	return req, nil
}
