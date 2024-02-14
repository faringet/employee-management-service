package transport

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/engagerocketco/go-common/auth0"
	"github.com/engagerocketco/go-common/consts"
	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"go.uber.org/zap"
)

func getEmail(ctx context.Context) (string, error) {
	var (
		email string
		err   error
	)
	if email, err = auth0.GetEmail(ctx); err != nil {
		switch {
		case errors.Is(err, auth0.NoValidatedClaimsErr):
			return "", ie.Error{
				Code:    http.StatusInternalServerError,
				Message: "could not find claims",
			}
		case errors.Is(err, auth0.NoCustomClaimsErr):
			return "", ie.Error{
				Code:    http.StatusForbidden,
				Message: "user email not found",
			}
		default:
			return "", ie.Error{
				Code:    http.StatusInternalServerError,
				Message: "unknown error from claims validation",
			}
		}
	}

	return email, nil
}

func parsePagination(r *http.Request, req *endpoints.PaginationRequest) error {
	if sortBy := r.URL.Query().Get("sort_by"); sortBy != "" {
		req.SortBy = sortBy
	}
	if sortType := r.URL.Query().Get("sort_type"); sortType != "" {
		req.SortType = sortType
	}
	if searchBy := r.URL.Query().Get("search_by"); searchBy != "" {
		req.SearchBy = searchBy
	}

	searchValue := r.URL.Query().Get("search_value")

	limitStr := r.URL.Query().Get("limit")

	offsetStr := r.URL.Query().Get("offset")

	var errDetails []ie.Details

	if req.SearchBy != "" && searchValue != "" {
		req.SearchValue = searchValue
	} else if req.SearchBy != "" && searchValue == "" {
		transportLogger.Warn("incorrect search_value format", zap.String("value", searchValue))
		errDetails = append(errDetails, ie.Details{
			Field:   "search_value",
			Message: "missing search_value parameter",
		})
	}

	if len(limitStr) > 0 {
		var err error
		req.Limit, err = strconv.Atoi(limitStr)

		if err != nil {
			transportLogger.Warn("incorrect limit format", zap.Error(err))
			errDetails = append(errDetails, ie.Details{
				Field:   "limit",
				Message: "limit value must be an integer",
			})
		}
	} else {
		req.Limit = consts.GetRequestLimit
	}

	if len(offsetStr) > 0 {
		var err error
		req.Offset, err = strconv.Atoi(offsetStr)

		if err != nil {
			transportLogger.Warn("incorrect offset format", zap.Error(err))
			errDetails = append(errDetails, ie.Details{
				Field:   "offset",
				Message: "offset value must be an integer",
			})
		}
	} else {
		req.Offset = consts.GetRequestOffset
	}

	if len(errDetails) > 0 {
		return ie.RequestValidationFailed(errDetails)
	}

	return nil
}
