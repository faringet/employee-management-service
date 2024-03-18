package repomodel

import (
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
)

type GetSmProject struct {
	model.SmProject
	SmProjectType *model.SmProjectType 
	Attributes *model.Attributes 
	SmAttributeTriggers     []GetSmAttributeTriggers
	SmSurvey     []GetSmSurvey

}

type GetSmProjectRequest struct {
	PaginationRequest
}
