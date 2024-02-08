package transport

import (
	"context"
	"errors"
	"net/http"

	"github.com/engagerocketco/go-common/auth0"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
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
