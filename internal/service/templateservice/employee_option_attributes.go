package templateservice

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"go.uber.org/zap"
)

type EmployeeOptionAttributes interface {
	CreateEmployeeOptionAttributes(ctx context.Context, input BaseEmployeeOptionAttributesDTO) (*BaseIdResponse, error)
	CreateEmployeeOptionAttributesRange(ctx context.Context, input []BaseEmployeeOptionAttributesDTO) error
	DeleteEmployeeOptionAttributesByID(ctx context.Context, input BaseIdRequest) error
	GetEmployeeOptionAttributesByID(ctx context.Context, req BaseIdRequest) (*BaseEmployeeOptionAttributesDTO, error)
	GetEmployeeOptionAttributes(ctx context.Context, req GetEmployeeOptionAttributesRequest) (*GetEmployeeOptionAttributesResponse, error)
	UpdateEmployeeOptionAttributesByID(ctx context.Context, input BaseEmployeeOptionAttributesDTO) (*BaseIdResponse, error)
}

type BaseEmployeeOptionAttributesDTO struct {
	ModifierEmail string
	ID            int32              `json:"id"`
	Value         *string            `json:"value"`
	CreatedAt     *time.Time         `json:"created_at"`
	UpdatedAt     *time.Time         `json:"updated_at"`
	AttributeID   *int32             `json:"attribute_id"`
	EmployeeID    *int32             `json:"employee_id"`
	Attributes    *BaseAttributesDTO `json:"attributes"`
	Employees     *BaseEmployeesDTO  `json:"employees"`
}

type GetEmployeeOptionAttributesRequest struct {
	PaginationRequest
}

type CreateEmployeeOptionAttributesRangeDTO struct {
	CreatorEmail             string
	EmployeeOptionAttributes []BaseEmployeeOptionAttributesDTO
}

type GetEmployeeOptionAttributesResponse struct {
	PaginationResponse
	Data []BaseEmployeeOptionAttributesDTO `json:"data"`
}

func (s *service) CreateEmployeeOptionAttributes(ctx context.Context, input BaseEmployeeOptionAttributesDTO) (*BaseIdResponse, error) {
	accId := int32(1)

	req := model.EmployeeOptionAttributes{
		CreatedBy:   &accId,
		UpdatedBy:   &accId,
		Value:       input.Value,
		AttributeID: input.AttributeID,
		EmployeeID:  input.EmployeeID,
	}

	res, err := s.repo.CreateEmployeeOptionAttributes(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateEmployeeOptionAttributesRange(ctx context.Context, input []BaseEmployeeOptionAttributesDTO) error {
	var accId int32 = 1

	var rows []model.EmployeeOptionAttributes

	for _, row := range input {
		rows = append(rows, model.EmployeeOptionAttributes{
			CreatedBy:   &accId,
			UpdatedBy:   &accId,
			Value:       row.Value,
			AttributeID: row.AttributeID,
			EmployeeID:  row.EmployeeID,
		})
	}

	err := s.repo.CreateEmployeeOptionAttributesRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetEmployeeOptionAttributesByID(ctx context.Context, req BaseIdRequest) (*BaseEmployeeOptionAttributesDTO, error) {
	row, err := s.repo.GetEmployeeOptionAttributesByID(ctx, repomodel.BaseIdRequest{
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
			return nil, ie.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}

	}

	tq := &BaseEmployeeOptionAttributesDTO{
		ID:          row.ID,
		Value:       row.Value,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		AttributeID: row.AttributeID,
		EmployeeID:  row.EmployeeID,
		Attributes: &BaseAttributesDTO{
			ID:        row.Attributes.ID,
			Name:      row.Attributes.Name,
			IsDate:    row.Attributes.IsDate,
			CreatedAt: row.Attributes.CreatedAt,
			UpdatedAt: row.Attributes.UpdatedAt,
			CreatedBy: row.Attributes.CreatedBy,
			UpdatedBy: row.Attributes.UpdatedBy,
		},
		Employees: &BaseEmployeesDTO{
			ID:            row.Employees.ID,
			Email:         row.Employees.Email,
			PreferredName: row.Employees.PreferredName,
			FullName:      row.Employees.FullName,
			UniqueID:      row.Employees.UniqueID,
			UpdatedAt:     row.Employees.UpdatedAt,
			CreatedAt:     row.Employees.CreatedAt,
			UpdatedBy:     row.Employees.UpdatedBy,
			CreatedBy:     row.Employees.CreatedBy,
			ManagerEmail:  row.Employees.ManagerEmail,
		},
	}

	return tq, nil
}

func (s *service) GetEmployeeOptionAttributes(ctx context.Context, req GetEmployeeOptionAttributesRequest) (*GetEmployeeOptionAttributesResponse, error) {

	input := repomodel.GetEmployeeOptionAttributesRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetEmployeeOptionAttributes(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, ie.ErrDBRecordNotFound):
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			s.logger.Error("get all template questionaries", zap.Error(err))
			return nil, ie.ErrInternalServerError
		}
	}

	res := []BaseEmployeeOptionAttributesDTO{}
	for _, row := range rows {
		tq := BaseEmployeeOptionAttributesDTO{
			ID:          row.ID,
			Value:       row.Value,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
			AttributeID: row.AttributeID,
			EmployeeID:  row.EmployeeID,
			Attributes: &BaseAttributesDTO{
				ID:        row.Attributes.ID,
				Name:      row.Attributes.Name,
				IsDate:    row.Attributes.IsDate,
				CreatedAt: row.Attributes.CreatedAt,
				UpdatedAt: row.Attributes.UpdatedAt,
				CreatedBy: row.Attributes.CreatedBy,
				UpdatedBy: row.Attributes.UpdatedBy,
			},
			Employees: &BaseEmployeesDTO{
				ID:            row.Employees.ID,
				Email:         row.Employees.Email,
				PreferredName: row.Employees.PreferredName,
				FullName:      row.Employees.FullName,
				UniqueID:      row.Employees.UniqueID,
				UpdatedAt:     row.Employees.UpdatedAt,
				CreatedAt:     row.Employees.CreatedAt,
				UpdatedBy:     row.Employees.UpdatedBy,
				CreatedBy:     row.Employees.CreatedBy,
				ManagerEmail:  row.Employees.ManagerEmail,
			},
		}

		res = append(res, tq)
	}

	count, err := s.repo.CountEmployeeOptionAttributes(ctx)
	if err != nil {
		s.logger.Error("count EmployeeOptionAttributes", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetEmployeeOptionAttributesResponse{
		Data: res,
		PaginationResponse: PaginationResponse{
			ActivePage:   activePage,
			TotalCount:   *count,
			CountPerPage: req.Limit,
			TotalPages:   totalPages,
		},
	}
	return page_res, nil
}

func (s *service) UpdateEmployeeOptionAttributesByID(ctx context.Context, input BaseEmployeeOptionAttributesDTO) (*BaseIdResponse, error) {
	var accId int32 = 1

	res, err := s.repo.UpdateEmployeeOptionAttributesByID(ctx, model.EmployeeOptionAttributes{
		ID:          input.ID,
		UpdatedBy:   &accId,
		Value:       input.Value,
		AttributeID: input.AttributeID,
		EmployeeID:  input.EmployeeID,
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

func (s *service) DeleteEmployeeOptionAttributesByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteEmployeeOptionAttributesByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
