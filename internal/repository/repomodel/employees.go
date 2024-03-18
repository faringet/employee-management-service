package repomodel

import (
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
)

type GetEmployees struct {
	model.Employees
	EmployeeOptionAttributes     []GetEmployeeOptionAttributes

}

type GetEmployeesRequest struct {
	PaginationRequest
}
