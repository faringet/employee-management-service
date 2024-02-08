package transport

import (
	"context"
	"encoding/json"
	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func MakeCommunicationUpdateTemplateByIDHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeUpdateTemplateCommsEndpoint(s),
		updateCommunicationTemplateByIDDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}
func updateCommunicationTemplateByIDDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		transportLogger.Error("missing customer status id in path")
		return nil, ie.Error{
			Code:    http.StatusBadRequest,
			Message: "missing templates id in path",
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		transportLogger.Error("customer id must be an integer")
		return nil, ie.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Request validation failed",
			Details: []ie.Details{
				{
					Field:   "id",
					Message: "customer id must be an integer",
				},
			},
		}
	}

	var req endpoints.UpdateCommunicationTemplateByIDRequest
	req.ID = id

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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
			Message: "Incorrect request format",
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
