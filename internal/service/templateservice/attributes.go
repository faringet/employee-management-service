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

type Attributes interface {
	CreateAttributes(ctx context.Context, input BaseAttributesDTO) (*BaseIdResponse, error)
	CreateAttributesRange(ctx context.Context, input []BaseAttributesDTO) error
	DeleteAttributesByID(ctx context.Context, input BaseIdRequest) error
	GetAttributesByID(ctx context.Context, req BaseIdRequest) (*BaseAttributesDTO, error)
	GetAttributes(ctx context.Context, req GetAttributesRequest) (*GetAttributesResponse, error)
	UpdateAttributesByID(ctx context.Context, input BaseAttributesDTO) (*BaseIdResponse, error)
}

type BaseAttributesDTO struct {
	ModifierEmail string
	ID            int32      `json:"id"`
	Name          *string    `json:"name"`
	IsDate        *bool      `json:"is_date"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	CreatedBy     *int32     `json:"created_by"`
	UpdatedBy     *int32     `json:"updated_by"`
}

type GetAttributesRequest struct {
	PaginationRequest
}

type CreateAttributesRangeDTO struct {
	CreatorEmail string
	Attributes   []BaseAttributesDTO
}

type GetAttributesResponse struct {
	PaginationResponse
	Data []BaseAttributesDTO `json:"data"`
}

func (s *service) CreateAttributes(ctx context.Context, input BaseAttributesDTO) (*BaseIdResponse, error) {
	accId := int32(1)

	req := model.Attributes{
		CreatedBy: &accId,
		UpdatedBy: &accId,
		Name:      input.Name,
		IsDate:    input.IsDate,
	}

	res, err := s.repo.CreateAttributes(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateAttributesRange(ctx context.Context, input []BaseAttributesDTO) error {
	var rows []model.Attributes
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.Attributes{
			CreatedBy: &accId,
			UpdatedBy: &accId,
			Name:      row.Name,
			IsDate:    row.IsDate,
		})
	}

	err := s.repo.CreateAttributesRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetAttributesByID(ctx context.Context, req BaseIdRequest) (*BaseAttributesDTO, error) {
	row, err := s.repo.GetAttributesByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseAttributesDTO{
		ID:        row.ID,
		Name:      row.Name,
		IsDate:    row.IsDate,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		CreatedBy: row.CreatedBy,
		UpdatedBy: row.UpdatedBy,
	}

	return tq, nil
}

func (s *service) GetAttributes(ctx context.Context, req GetAttributesRequest) (*GetAttributesResponse, error) {

	input := repomodel.GetAttributesRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetAttributes(ctx, input)
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

	res := []BaseAttributesDTO{}
	for _, row := range rows {
		tq := BaseAttributesDTO{
			ID:        row.ID,
			Name:      row.Name,
			IsDate:    row.IsDate,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
			CreatedBy: row.CreatedBy,
			UpdatedBy: row.UpdatedBy,
		}

		res = append(res, tq)
	}

	count, err := s.repo.CountAttributes(ctx)
	if err != nil {
		s.logger.Error("count Attributes", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetAttributesResponse{
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

func (s *service) UpdateAttributesByID(ctx context.Context, input BaseAttributesDTO) (*BaseIdResponse, error) {
	var accId int32 = 1

	var now = time.Now().UTC()

	res, err := s.repo.UpdateAttributesByID(ctx, model.Attributes{
		ID:        int32(input.ID),
		UpdatedBy: &accId,
		UpdatedAt: &now,

		Name:   input.Name,
		IsDate: input.IsDate,
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

func (s *service) DeleteAttributesByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteAttributesByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
