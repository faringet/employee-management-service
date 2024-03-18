package repomodel

import "github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"

type GetSmSurveyRecepients struct {
	model.SmSurveyRecepients
	
	Employees *model.Employees 
	SmSurvey *model.SmSurvey 
	

}

type GetSmSurveyRecepientsRequest struct {
	PaginationRequest
}
