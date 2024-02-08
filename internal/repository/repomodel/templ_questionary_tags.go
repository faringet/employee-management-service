package repomodel

import "github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"

type GetTemplQuestionaryTags struct {
	model.TemplQuestionaryTags
	SurveyTags model.SurveyTags
}

type GetTemplQuestionaryTagsRequest struct {
	PaginationRequest
}
