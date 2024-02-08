package endpoints

import "github.com/engagerocketco/templates-api-svc/pkg/errors"

type EmptyResponseSample struct{}

type SwaggerSimpleError struct {
	Error string `json:"error"`
}

type SwaggerError struct {
	Error   string           `json:"error"`
	Details []errors.Details `json:"details"`
}
