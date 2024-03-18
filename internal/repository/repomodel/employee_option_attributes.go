package repomodel

import "github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"

type GetEmployeeOptionAttributes struct {
	model.EmployeeOptionAttributes
	
	Attributes *model.Attributes 
	Employees *model.Employees 
	

}

type GetEmployeeOptionAttributesRequest struct {
	PaginationRequest
}
