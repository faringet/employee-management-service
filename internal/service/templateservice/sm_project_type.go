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

type SmProjectType interface {
	CreateSmProjectType(ctx context.Context, input BaseSmProjectTypeDTO) (*BaseIdResponse, error)
	CreateSmProjectTypeRange(ctx context.Context, input []BaseSmProjectTypeDTO) error
	DeleteSmProjectTypeByID(ctx context.Context, input BaseIdRequest) error
	GetSmProjectTypeByID(ctx context.Context, req BaseIdRequest) (*BaseSmProjectTypeDTO, error)
	GetSmProjectType(ctx context.Context, req GetSmProjectTypeRequest) (*GetSmProjectTypeResponse, error)
	UpdateSmProjectTypeByID(ctx context.Context, input BaseSmProjectTypeDTO) (*BaseIdResponse, error)
}

type BaseSmProjectTypeDTO struct {
	ModifierEmail string
	ID            int32      `json:"id"`
	Name          *string    `json:"name"`
	Code          *string    `json:"code"`
	Description   *string    `json:"description"`
	UpdatedAt     *time.Time `json:"updated_at"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedBy     *int32     `json:"updated_by"`
	CreatedBy     *int32     `json:"created_by"`
}

type GetSmProjectTypeRequest struct {
	PaginationRequest
}

type CreateSmProjectTypeRangeDTO struct {
	CreatorEmail  string
	SmProjectType []BaseSmProjectTypeDTO
}

type GetSmProjectTypeResponse struct {
	PaginationResponse
	Data []BaseSmProjectTypeDTO `json:"data"`
}

func (s *service) CreateSmProjectType(ctx context.Context, input BaseSmProjectTypeDTO) (*BaseIdResponse, error) {

	accId := int32(1)

	req := model.SmProjectType{
		CreatedBy:   &accId,
		UpdatedBy:   &accId,
		Name:        input.Name,
		Code:        input.Code,
		Description: input.Description,
	}

	res, err := s.repo.CreateSmProjectType(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateSmProjectTypeRange(ctx context.Context, input []BaseSmProjectTypeDTO) error {

	var rows []model.SmProjectType
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.SmProjectType{
			CreatedBy:   &accId,
			UpdatedBy:   &accId,
			Name:        row.Name,
			Code:        row.Code,
			Description: row.Description,
		})
	}

	err := s.repo.CreateSmProjectTypeRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetSmProjectTypeByID(ctx context.Context, req BaseIdRequest) (*BaseSmProjectTypeDTO, error) {
	row, err := s.repo.GetSmProjectTypeByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseSmProjectTypeDTO{
		ID:          row.ID,
		Name:        row.Name,
		Code:        row.Code,
		Description: row.Description,
		UpdatedAt:   row.UpdatedAt,
		CreatedAt:   row.CreatedAt,
		UpdatedBy:   row.UpdatedBy,
		CreatedBy:   row.CreatedBy,
	}

	return tq, nil
}

func (s *service) GetSmProjectType(ctx context.Context, req GetSmProjectTypeRequest) (*GetSmProjectTypeResponse, error) {

	input := repomodel.GetSmProjectTypeRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetSmProjectType(ctx, input)
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

	res := []BaseSmProjectTypeDTO{}
	for _, row := range rows {
		tq := BaseSmProjectTypeDTO{
			ID:          row.ID,
			Name:        row.Name,
			Code:        row.Code,
			Description: row.Description,
			UpdatedAt:   row.UpdatedAt,
			CreatedAt:   row.CreatedAt,
			UpdatedBy:   row.UpdatedBy,
			CreatedBy:   row.CreatedBy,
		}

		res = append(res, tq)
	}

	count, err := s.repo.CountSmProjectType(ctx)
	if err != nil {
		s.logger.Error("count SmProjectType", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetSmProjectTypeResponse{
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

func (s *service) UpdateSmProjectTypeByID(ctx context.Context, input BaseSmProjectTypeDTO) (*BaseIdResponse, error) {

	var accId int32 = 1

	var now = time.Now().UTC()

	res, err := s.repo.UpdateSmProjectTypeByID(ctx, model.SmProjectType{
		ID:        int32(input.ID),
		UpdatedBy: &accId,
		UpdatedAt: &now,

		Name:        input.Name,
		Code:        input.Code,
		Description: input.Description,
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

func (s *service) DeleteSmProjectTypeByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteSmProjectTypeByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
