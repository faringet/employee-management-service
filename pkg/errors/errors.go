package errors

import (
	"errors"
	"net/http"
)

type Details struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Code    int       `json:"-"`
	Message string    `json:"error"`
	Details []Details `json:"details,omitempty"`
}

func (err Error) Error() string {
	return err.Message
}

var ErrDBRecordNotFound = errors.New("db record not found")

var ErrInternalServerError error = Error{
	Code:    http.StatusInternalServerError,
	Message: "internal server error",
}

func RequestValidationFailed(details []Details) Error {
	return Error{
		Details: details,
		Code:    http.StatusUnprocessableEntity,
		Message: "Request validation failed",
	}
}
