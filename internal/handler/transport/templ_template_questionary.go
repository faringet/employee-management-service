package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/engagerocketco/templates-api-svc/pkg/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func MakeGetTemplTemplateQuestionaryByIDHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeGetTemplTemplateQuestionaryByIDEndpoint(s),
		getTemplTemplateQuestionaryByIDDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func getTemplTemplateQuestionaryByIDDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		transportLogger.Warn("missing survey_tags id in path")
		return nil, errors.Error{
			Code:    http.StatusBadRequest,
			Message: "missing survey_tags id in path",
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		transportLogger.Warn("survey_tags id must be an integer")
		return nil, errors.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "request validation failed",
			Details: []errors.Details{
				{
					Field:   "id",
					Message: "survey_tags id must be an integer",
				},
			},
		}
	}

	return endpoints.BaseIdRequest{
		ID: id,
	}, nil
}

func MakeGetTemplTemplateQuestionaryHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeGetTemplTemplateQuestionaryEndpoint(s),
		getTemplTemplateQuestionaryDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func getTemplTemplateQuestionaryDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	sortBy := r.URL.Query().Get("sort_by")
	sortType := r.URL.Query().Get("sort_type")
	searchBy := r.URL.Query().Get("search_by")
	searchValue := r.URL.Query().Get("search_value")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	var req endpoints.GetTemplTemplateQuestionaryRequest
	validationError := errors.Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Request validation failed",
	}
	var errDetails []errors.Details

	req.SortBy = sortBy
	req.SortType = sortType
	req.SearchBy = searchBy

	if len(req.SearchBy) > 0 && len(searchValue) > 0 {
		req.SearchValue = searchValue
	} else if len(req.SearchBy) > 0 && len(searchValue) == 0 {
		transportLogger.Warn("incorrect search_value format", zap.String("value", searchValue))
		errDetails = append(errDetails, errors.Details{
			Field:   "search_value",
			Message: "missing search_value parameter",
		})
	}

	if len(limitStr) > 0 {
		var err error
		req.Limit, err = strconv.Atoi(limitStr)

		if err != nil {
			transportLogger.Warn("incorrect limit format", zap.Error(err))
			errDetails = append(errDetails, errors.Details{
				Field:   "limit",
				Message: "limit value must be an integer",
			})
		}
	} else {
		req.Limit = 500 //TODO
	}

	if len(offsetStr) > 0 {
		var err error
		req.Offset, err = strconv.Atoi(offsetStr)

		if err != nil {
			transportLogger.Warn("incorrect offset format", zap.Error(err))
			errDetails = append(errDetails, errors.Details{
				Field:   "offset",
				Message: "offset value must be an integer",
			})
		}
	} else {
		req.Offset = 0 //TODO
	}

	if len(errDetails) > 0 {
		validationError.Details = errDetails
		return nil, validationError
	}

	return req, nil
}

func MakeCreateTemplTemplateQuestionaryHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeCreateTemplTemplateQuestionaryEndpoint(s),
		createTemplTemplateQuestionaryDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func createTemplTemplateQuestionaryDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	email, err := getEmail(ctx)
	if err != nil {
		return nil, err
	}

	var req endpoints.CreateTemplTemplateQuestionaryRequest
	req.CreatorEmail = email

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		transportLogger.Warn("decode the request body", zap.Error(err))
		return nil, errors.Error{
			Code:    http.StatusBadRequest,
			Message: "unable to decode the request body",
		}
	}

	requestValidator := validator.New()
	err = requestValidator.Struct(req)
	if err != nil {
		transportLogger.Warn("request body is not valid", zap.Error(err))
		errToReturn := errors.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "incorrect request format",
		}
		for _, err := range err.(validator.ValidationErrors) {
			details := errors.Details{
				Field:   err.Field(),
				Message: err.Error(),
			}
			errToReturn.Details = append(errToReturn.Details, details)
		}
		return nil, errToReturn
	}

	return req, nil
}

func MakeCreateTemplTemplateQuestionaryRangeHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeCreateTemplTemplateQuestionaryRangeEndpoint(s),
		createTemplTemplateQuestionaryRangeDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func createTemplTemplateQuestionaryRangeDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	email, err := getEmail(ctx)
	if err != nil {
		return nil, err
	}

	var req endpoints.CreateTemplTemplateQuestionaryRangeRequest
	req.CreatorEmail = email

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		transportLogger.Warn("decode the request body", zap.Error(err))
		return nil, errors.Error{
			Code:    http.StatusBadRequest,
			Message: "unable to decode the request body",
		}
	}

	requestValidator := validator.New()
	err = requestValidator.Struct(req)
	if err != nil {
		transportLogger.Warn("request body is not valid", zap.Error(err))
		errToReturn := errors.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "incorrect request format",
		}
		for _, err := range err.(validator.ValidationErrors) {
			details := errors.Details{
				Field:   err.Field(),
				Message: err.Error(),
			}
			errToReturn.Details = append(errToReturn.Details, details)
		}
		return nil, errToReturn
	}

	return req, nil
}

func MakeDeleteTemplTemplateQuestionaryByIDHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeDeleteTemplTemplateQuestionaryByIDEndpoint(s),
		deleteTemplTemplateQuestionaryDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func deleteTemplTemplateQuestionaryDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		transportLogger.Warn("missing status id in path")
		return nil, errors.Error{
			Code:    http.StatusBadRequest,
			Message: "missing status id in path",
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		transportLogger.Warn("template questionaries id must be an integer")
		return nil, errors.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "request validation failed",
			Details: []errors.Details{
				{
					Field:   "id",
					Message: "template questionaries id must be an integer",
				},
			},
		}
	}

	return endpoints.BaseIdRequest{
		ID: id,
	}, nil
}

func MakeUpdateTemplTemplateQuestionaryByIDHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeUpdateTemplTemplateQuestionaryEndpoint(s),
		updateTemplTemplateQuestionaryTemplateByIDDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}
func updateTemplTemplateQuestionaryTemplateByIDDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	email, err := getEmail(ctx)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		transportLogger.Error("missing survey_tags status id in path")
		return nil, errors.Error{
			Code:    http.StatusBadRequest,
			Message: "missing templates id in path",
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		transportLogger.Error("survey_tags id must be an integer")
		return nil, errors.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Request validation failed",
			Details: []errors.Details{
				{
					Field:   "id",
					Message: "survey_tags id must be an integer",
				},
			},
		}
	}

	var req endpoints.UpdateTemplTemplateQuestionaryRequest
	req.ID = int32(id)
	req.UpdaterEmail = email

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		transportLogger.Error("decode the request body", zap.Error(err))
		return nil, errors.Error{
			Code:    http.StatusBadRequest,
			Message: "unable to decode the request body",
		}
	}

	requestValidator := validator.New()
	err = requestValidator.Struct(req)

	if err != nil {
		transportLogger.Error("request body is not valid", zap.Error(err))
		errToReturn := errors.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Incorrect request format",
		}
		for _, err := range err.(validator.ValidationErrors) {
			details := errors.Details{
				Field:   err.Field(),
				Message: err.Error(),
			}
			errToReturn.Details = append(errToReturn.Details, details)
		}
		return nil, errToReturn
	}

	return req, nil
}
