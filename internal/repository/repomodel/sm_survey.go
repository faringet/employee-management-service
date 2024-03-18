package repomodel

import "github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"

type GetSmSurvey struct {
	model.SmSurvey

	SmSurveyStatus *model.SmSurveyStatus
	SmProject      *model.SmProject
}

type GetSmSurveyRequest struct {
	PaginationRequest
}
