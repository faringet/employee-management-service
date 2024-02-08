package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// "github.com/engagerocketco/go-common/consts"
	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/engagerocketco/templates-api-svc/pkg/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func MakeGetSurveyTagsByIDHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeGetSurveyTagsByIDEndpoint(s),
		getSurveyTagsByIDDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func getSurveyTagsByIDDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func MakeGetSurveyTagsHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeGetSurveyTagsEndpoint(s),
		getSurveyTagsDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func getSurveyTagsDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoints.GetSurveyTagsRequest{}, nil
}

func MakeGetSurveyTagsPaginationHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeGetSurveyTagsPaginationEndpoint(s),
		getSurveyTagsPaginationDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func getSurveyTagsPaginationDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// return endpoints.GetSurveyTagsRequest{}, nil

	sortBy := r.URL.Query().Get("sort_by")
	sortType := r.URL.Query().Get("sort_type")
	searchBy := r.URL.Query().Get("search_by")
	searchValue := r.URL.Query().Get("search_value")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	log.Println(limitStr)

	var req endpoints.GetSurveyTagsRequest
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

func MakeCreateSurveyTagsHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeCreateSurveyTagsEndpoint(s),
		createSurveyTagsDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func createSurveyTagsDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	email, err := getEmail(ctx)
	if err != nil {
		return nil, err
	}

	var req endpoints.CreateSurveyTagsRequest
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

func MakeCreateSurveyTagsRangeHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeCreateSurveyTagsRangeEndpoint(s),
		createSurveyTagsRangeDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func createSurveyTagsRangeDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	email, err := getEmail(ctx)
	if err != nil {
		return nil, err
	}

	var req endpoints.CreateSurveyTagsRangeRequest
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

func MakeDeleteSurveyTagsByIDHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeDeleteSurveyTagsByIDEndpoint(s),
		deleteSurveyTagsDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}

func deleteSurveyTagsDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func MakeUpdateSurveyTagsByIDHandler(s templateservice.Service, logger *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerFinalizer(finalyzer),
		kithttp.ServerErrorEncoder(encodeError),
	}

	h := kithttp.NewServer(
		endpoints.MakeUpdateSurveyTagsEndpoint(s),
		updateSurveyTagsTemplateByIDDecodeRequest,
		encodeResponse,
		opts...,
	)

	return h
}
func updateSurveyTagsTemplateByIDDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

	var req endpoints.UpdateSurveyTagsRequest
	req.ID = int32(id)

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
