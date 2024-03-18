package repomodel

import "github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"

type GetSmAttributeTriggers struct {
	model.SmAttributeTriggers
	
	Attributes *model.Attributes 
	SmProject *model.SmProject 
	

}

type GetSmAttributeTriggersRequest struct {
	PaginationRequest
}
