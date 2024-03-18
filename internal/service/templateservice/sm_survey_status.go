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

type SmSurveyStatus interface {
	CreateSmSurveyStatus(ctx context.Context, input BaseSmSurveyStatusDTO) (*BaseIdResponse, error)
	CreateSmSurveyStatusRange(ctx context.Context, input []BaseSmSurveyStatusDTO) error
	DeleteSmSurveyStatusByID(ctx context.Context, input BaseIdRequest) error
	GetSmSurveyStatusByID(ctx context.Context, req BaseIdRequest) (*BaseSmSurveyStatusDTO, error)
	GetSmSurveyStatus(ctx context.Context, req GetSmSurveyStatusRequest) (*GetSmSurveyStatusResponse, error)
	UpdateSmSurveyStatusByID(ctx context.Context, input BaseSmSurveyStatusDTO) (*BaseIdResponse, error)
}

type BaseSmSurveyStatusDTO struct {
	ModifierEmail string
	ID            int32      `json:"id"`
	Name          *string    `json:"name"`
	Code          *string    `json:"code"`
	Description   *string    `json:"description"`
	UpdatedAt     *time.Time `json:"updated_at"`
	CreatedBy     *int32     `json:"created_by"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedBy     *int32     `json:"updated_by"`
}

type GetSmSurveyStatusRequest struct {
	PaginationRequest
}

type CreateSmSurveyStatusRangeDTO struct {
	CreatorEmail   string
	SmSurveyStatus []BaseSmSurveyStatusDTO
}

type GetSmSurveyStatusResponse struct {
	PaginationResponse
	Data []BaseSmSurveyStatusDTO `json:"data"`
}

func (s *service) CreateSmSurveyStatus(ctx context.Context, input BaseSmSurveyStatusDTO) (*BaseIdResponse, error) {

	accId := int32(1)

	req := model.SmSurveyStatus{
		CreatedBy:   &accId,
		UpdatedBy:   &accId,
		Name:        input.Name,
		Code:        input.Code,
		Description: input.Description,
	}

	res, err := s.repo.CreateSmSurveyStatus(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateSmSurveyStatusRange(ctx context.Context, input []BaseSmSurveyStatusDTO) error {

	var rows []model.SmSurveyStatus
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.SmSurveyStatus{
			CreatedBy:   &accId,
			UpdatedBy:   &accId,
			Name:        row.Name,
			Code:        row.Code,
			Description: row.Description,
		})
	}

	err := s.repo.CreateSmSurveyStatusRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetSmSurveyStatusByID(ctx context.Context, req BaseIdRequest) (*BaseSmSurveyStatusDTO, error) {
	row, err := s.repo.GetSmSurveyStatusByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseSmSurveyStatusDTO{
		ID:          row.ID,
		Name:        row.Name,
		Code:        row.Code,
		Description: row.Description,
		UpdatedAt:   row.UpdatedAt,
		CreatedBy:   row.CreatedBy,
		CreatedAt:   row.CreatedAt,
		UpdatedBy:   row.UpdatedBy,
	}

	return tq, nil
}

func (s *service) GetSmSurveyStatus(ctx context.Context, req GetSmSurveyStatusRequest) (*GetSmSurveyStatusResponse, error) {

	input := repomodel.GetSmSurveyStatusRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetSmSurveyStatus(ctx, input)
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

	res := []BaseSmSurveyStatusDTO{}
	for _, row := range rows {
		tq := BaseSmSurveyStatusDTO{
			ID:          row.ID,
			Name:        row.Name,
			Code:        row.Code,
			Description: row.Description,
			UpdatedAt:   row.UpdatedAt,
			CreatedBy:   row.CreatedBy,
			CreatedAt:   row.CreatedAt,
			UpdatedBy:   row.UpdatedBy,
		}

		res = append(res, tq)
	}

	count, err := s.repo.CountSmSurveyStatus(ctx)
	if err != nil {
		s.logger.Error("count SmSurveyStatus", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetSmSurveyStatusResponse{
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

func (s *service) UpdateSmSurveyStatusByID(ctx context.Context, input BaseSmSurveyStatusDTO) (*BaseIdResponse, error) {

	var accId int32 = 1

	var now = time.Now().UTC()

	res, err := s.repo.UpdateSmSurveyStatusByID(ctx, model.SmSurveyStatus{
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

func (s *service) DeleteSmSurveyStatusByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteSmSurveyStatusByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
