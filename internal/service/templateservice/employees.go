package templateservice

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"go.uber.org/zap"
)

type Employees interface {
	CreateEmployees(ctx context.Context, input BaseEmployeesDTO) (*BaseIdResponse, error)
	CreateEmployeesRange(ctx context.Context, input []BaseEmployeesDTO) error
	DeleteEmployeesByID(ctx context.Context, input BaseIdRequest) error
	GetEmployeesByID(ctx context.Context, req BaseIdRequest) (*BaseEmployeesDTO, error)
	GetEmployees(ctx context.Context, req GetEmployeesRequest) (*GetEmployeesResponse, error)
	UpdateEmployeesByID(ctx context.Context, input BaseEmployeesDTO) (*BaseIdResponse, error)
}

type BaseEmployeesDTO struct {
	ModifierEmail string
	ID            int32      `json:"id"`
	Email         *string    `json:"email"`
	PreferredName *string    `json:"preferred_name"`
	FullName      *string    `json:"full_name"`
	UniqueID      *string    `json:"unique_id"`
	UpdatedAt     *time.Time `json:"updated_at"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedBy     *int32     `json:"updated_by"`
	CreatedBy     *int32     `json:"created_by"`
	ManagerEmail  *string    `json:"manager_email"`

	EmployeeOptionAttributes []BaseEmployeeOptionAttributesDTO `json:"employee_option_attributes"`
}

type GetEmployeesRequest struct {
	PaginationRequest
}

type CreateEmployeesRangeDTO struct {
	CreatorEmail string
	Employees    []BaseEmployeesDTO
}

type GetEmployeesResponse struct {
	PaginationResponse
	Data []BaseEmployeesDTO `json:"data"`
}

func (s *service) CreateEmployees(ctx context.Context, input BaseEmployeesDTO) (*BaseIdResponse, error) {
	accId := int32(1)

	req := model.Employees{
		CreatedBy:     &accId,
		UpdatedBy:     &accId,
		Email:         input.Email,
		PreferredName: input.PreferredName,
		FullName:      input.FullName,
		UniqueID:      input.UniqueID,
		ManagerEmail:  input.ManagerEmail,
	}

	res, err := s.repo.CreateEmployees(ctx, req)
	if err != nil {
		s.logger.Error("create Employees", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateEmployeesRange(ctx context.Context, input []BaseEmployeesDTO) error {
	var rows []model.Employees
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.Employees{
			CreatedBy:     &accId,
			UpdatedBy:     &accId,
			Email:         row.Email,
			PreferredName: row.PreferredName,
			FullName:      row.FullName,
			UniqueID:      row.UniqueID,
			ManagerEmail:  row.ManagerEmail,
		})
	}

	err := s.repo.CreateEmployeesRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetEmployeesByID(ctx context.Context, req BaseIdRequest) (*BaseEmployeesDTO, error) {
	row, err := s.repo.GetEmployeesByID(ctx, repomodel.BaseIdRequest{
		ID: int32(req.ID),
	})
	if err != nil {
		switch {
		case errors.Is(err, ie.ErrDBRecordNotFound):
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			//  s.logger.Error(, zap.Error(err))
			return nil, ie.ErrInternalServerError
		}
	}

	tq := BaseEmployeesDTO{
		ID:            row.ID,
		UpdatedAt:     row.UpdatedAt,
		CreatedAt:     row.CreatedAt,
		UpdatedBy:     row.UpdatedBy,
		CreatedBy:     row.CreatedBy,
		Email:         row.Email,
		PreferredName: row.PreferredName,
		FullName:      row.FullName,
		UniqueID:      row.UniqueID,
		ManagerEmail:  row.ManagerEmail,
	}

	employee_option_attributes := []BaseEmployeeOptionAttributesDTO{}
	for _, r := range row.EmployeeOptionAttributes {
		employee_option_attributes = append(employee_option_attributes, BaseEmployeeOptionAttributesDTO{
			ID:          r.ID,
			Value:       r.Value,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
			AttributeID: r.AttributeID,
			EmployeeID:  r.EmployeeID,
		})
	}
	tq.EmployeeOptionAttributes = employee_option_attributes

	return &tq, nil
}

func (s *service) GetEmployees(ctx context.Context, req GetEmployeesRequest) (*GetEmployeesResponse, error) {

	input := repomodel.GetEmployeesRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	log.Println(input)

	rows, err := s.repo.GetEmployees(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, ie.ErrDBRecordNotFound):
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:

			//  s.logger.Error(, zap.Error(err))
			return nil, ie.ErrInternalServerError
		}
	}
	res := []BaseEmployeesDTO{}

	for _, row := range rows {
		tq := BaseEmployeesDTO{
			ID:            row.ID,
			UpdatedAt:     row.UpdatedAt,
			CreatedAt:     row.CreatedAt,
			UpdatedBy:     row.UpdatedBy,
			CreatedBy:     row.CreatedBy,
			Email:         row.Email,
			PreferredName: row.PreferredName,
			FullName:      row.FullName,
			UniqueID:      row.UniqueID,
			ManagerEmail:  row.ManagerEmail,
		}

		employee_option_attributes := []BaseEmployeeOptionAttributesDTO{}
		for _, r := range row.EmployeeOptionAttributes {
			employee_option_attributes = append(employee_option_attributes, BaseEmployeeOptionAttributesDTO{
				ID:          r.ID,
				Value:       r.Value,
				CreatedAt:   r.CreatedAt,
				UpdatedAt:   r.UpdatedAt,
				AttributeID: r.AttributeID,
				EmployeeID:  r.EmployeeID,
			})
		}
		tq.EmployeeOptionAttributes = employee_option_attributes

		res = append(res, tq)

	}

	count, err := s.repo.CountEmployees(ctx)
	if err != nil {
		s.logger.Error("count Employees", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetEmployeesResponse{
		Data: res,
		PaginationResponse: PaginationResponse{
			ActivePage:   activePage,
			TotalCount:   *count,
			CountPerPage: req.Limit,
			TotalPages:   totalPages,
		},
	}
	log.Println(page_res)

	return page_res, nil
}

func (s *service) UpdateEmployeesByID(ctx context.Context, input BaseEmployeesDTO) (*BaseIdResponse, error) {
	var accId int32 = 1

	res, err := s.repo.UpdateEmployeesByID(ctx, model.Employees{
		ID:        int32(input.ID),
		UpdatedBy: &accId,

		Email:         input.Email,
		PreferredName: input.PreferredName,
		FullName:      input.FullName,
		CreatedBy:     input.CreatedBy,
		ManagerEmail:  input.ManagerEmail,
	})
	if err != nil {
		switch {
		case errors.Is(err, ie.ErrDBRecordNotFound):
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			return nil, ie.ErrInternalServerError
		}
	}

	e := BaseIdResponse(*res)

	return &e, nil
}

func (s *service) DeleteEmployeesByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteEmployeesByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
