package repomodel

import (
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
)

type GetTemplTemplateQuestionary struct {
	model.TemplTemplateQuestionary
	TemplRecomendedFrequancy model.TemplRecomendedFrequancy
	TemplQuestionaryTags     []GetTemplQuestionaryTags
}

type GetTemplTemplateQuestionaryRequest struct {
	PaginationRequest
}
