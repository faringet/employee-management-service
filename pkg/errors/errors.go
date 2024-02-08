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

var ErrTemplateQuestionaryNotFound = errors.New("template questionary not found")
var ErrAccountNotFound = errors.New("account not found")
var ErrSurveyTagsNotFound = errors.New("survey_tags not found")
var ErrInternalServerError error = Error{
	Code:    http.StatusInternalServerError,
	Message: "internal server error",
}

var ErrCustom1 error = Error{
	Code:    http.StatusInternalServerError,
	Message: "internal server error1",
}
var ErrCustom2 error = Error{
	Code:    http.StatusInternalServerError,
	Message: "internal server error2",
}
var ErrCustom3 error = Error{
	Code:    http.StatusInternalServerError,
	Message: "internal server error3",
}
