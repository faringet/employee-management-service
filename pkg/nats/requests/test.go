package requests

import "errors"

const TestSubject = "entity.test"

type TestRequest struct {
	Text string `json:"text"`
}

type TestResponse struct {
	Response string `json:"response"`
}

var (
	TestErrSomethingWentWrong = errors.New("Something went wrong!")
	TestErrIncorrectData      = errors.New("Incorrect request data!")
)
