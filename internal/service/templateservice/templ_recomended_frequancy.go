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

type TemplRecomendedFrequancy interface {
	CreateTemplRecomendedFrequancy(ctx context.Context, input BaseTemplRecomendedFrequancyDTO) (*BaseIdResponse, error)
	CreateTemplRecomendedFrequancyRange(ctx context.Context, input []BaseTemplRecomendedFrequancyDTO) error
	DeleteTemplRecomendedFrequancyByID(ctx context.Context, input BaseIdRequest) error
	GetTemplRecomendedFrequancyByID(ctx context.Context, req BaseIdRequest) (*BaseTemplRecomendedFrequancyDTO, error)
	GetTemplRecomendedFrequancy(ctx context.Context, req GetTemplRecomendedFrequancyRequest) ([]BaseTemplRecomendedFrequancyDTO, error)
	UpdateTemplRecomendedFrequancyByID(ctx context.Context, input BaseTemplRecomendedFrequancyDTO) (*BaseIdResponse, error)
}

type BaseTemplRecomendedFrequancyDTO struct {
	ID              int32      `json:"id"`
	UpdatedAt       *time.Time `json:"updated_at"` //как в БД
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedBy       *int32     `json:"updated_by"`
	CreatedBy       *int32     `json:"created_by"`
	Name            *string    `json:"name"` //как в БД
	Description     *string    `json:"description"`
	Code            *string    `json:"code"`
	QueueNumber     *int32     `json:"queueNumber"`
	IconColor       *string    `json:"iconColor"`
	IdCustomSvgIcon *int32     `json:"idCustomSvgIcon"`
}

type GetTemplRecomendedFrequancyRequest struct {
	PaginationRequest
}

type CreateTemplRecomendedFrequancyRangeDTO struct {
	CreatorEmail             string
	TemplRecomendedFrequancy []BaseTemplRecomendedFrequancyDTO
}

func (s *service) CreateTemplRecomendedFrequancy(ctx context.Context, input BaseTemplRecomendedFrequancyDTO) (*BaseIdResponse, error) {
	// accInfo, err := s.natsService.GetAccountInfo(ctx, input.CreatorEmail)
	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, ie.ErrAccountNotFound):
	// 		return nil, ie.Error{
	// 			Code:    http.StatusNotFound,
	// 			Message: err.Error(),
	// 		}
	// 	default:
	// 		s.logger.Error("get acc info", zap.Error(err), zap.Any("create request", input))
	// 		return nil, ie.ErrInternalServerError
	// 	}
	// }
	accId := int32(1)

	req := model.TemplRecomendedFrequancy{
		CreatedBy:       &accId,
		UpdatedBy:       &accId,
		Name:            input.Name,
		Description:     input.Description,
		Code:            input.Code,
		QueueNumber:     input.QueueNumber,
		IconColor:       input.IconColor,
		IdCustomSvgIcon: input.IdCustomSvgIcon,
	}

	res, err := s.repo.CreateTemplRecomendedFrequancy(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateTemplRecomendedFrequancyRange(ctx context.Context, input []BaseTemplRecomendedFrequancyDTO) error {
	// accInfo, err := s.natsService.GetAccountInfo(ctx, input.CreatorEmail)
	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, ie.ErrAccountNotFound):
	// 		return nil, ie.Error{
	// 			Code:    http.StatusNotFound,
	// 			Message: err.Error(),
	// 		}
	// 	default:
	// 		s.logger.Error("get acc info", zap.Error(err), zap.Any("create request", input))
	// 		return nil, ie.ErrInternalServerError
	// 	}
	// }
	var rows []model.TemplRecomendedFrequancy
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.TemplRecomendedFrequancy{
			CreatedBy:       &accId,
			UpdatedBy:       &accId,
			Name:            row.Name,
			Description:     row.Description,
			Code:            row.Code,
			QueueNumber:     row.QueueNumber,
			IconColor:       row.IconColor,
			IdCustomSvgIcon: row.IdCustomSvgIcon,
		})
	}

	err := s.repo.CreateTemplRecomendedFrequancyRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetTemplRecomendedFrequancyByID(ctx context.Context, req BaseIdRequest) (*BaseTemplRecomendedFrequancyDTO, error) {
	row, err := s.repo.GetTemplRecomendedFrequancyByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseTemplRecomendedFrequancyDTO{
		ID:              row.ID,
		UpdatedAt:       row.UpdatedAt,
		CreatedAt:       row.CreatedAt,
		UpdatedBy:       row.UpdatedBy,
		CreatedBy:       row.CreatedBy,
		Name:            row.Name,
		Description:     row.Description,
		Code:            row.Code,
		QueueNumber:     row.QueueNumber,
		IconColor:       row.IconColor,
		IdCustomSvgIcon: row.IdCustomSvgIcon,
	}

	return tq, nil
}

func (s *service) GetTemplRecomendedFrequancy(ctx context.Context, req GetTemplRecomendedFrequancyRequest) ([]BaseTemplRecomendedFrequancyDTO, error) {

	input := repomodel.GetTemplRecomendedFrequancyRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetTemplRecomendedFrequancy(ctx, input)
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

	res := []BaseTemplRecomendedFrequancyDTO{}
	for _, row := range rows {
		tq := BaseTemplRecomendedFrequancyDTO{
			ID:              row.ID,
			UpdatedAt:       row.UpdatedAt,
			CreatedAt:       row.CreatedAt,
			UpdatedBy:       row.UpdatedBy,
			CreatedBy:       row.CreatedBy,
			Name:            row.Name,
			Description:     row.Description,
			Code:            row.Code,
			QueueNumber:     row.QueueNumber,
			IconColor:       row.IconColor,
			IdCustomSvgIcon: row.IdCustomSvgIcon,
		}

		res = append(res, tq)
	}

	return res, nil
}

func (s *service) UpdateTemplRecomendedFrequancyByID(ctx context.Context, input BaseTemplRecomendedFrequancyDTO) (*BaseIdResponse, error) {
	// accInfo, err := s.natsService.GetAccountInfo(ctx, input.UpdaterEmail)
	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, ie.ErrAccountNotFound):
	// 		return nil, ie.Error{
	// 			Code:    http.StatusNotFound,
	// 			Message: err.Error(),
	// 		}
	// 	default:
	// 		return nil, ie.ErrInternalServerError
	// 	}
	// }
	var accId int32 = 1

	var now = time.Now().UTC()

	res, err := s.repo.UpdateTemplRecomendedFrequancyByID(ctx, model.TemplRecomendedFrequancy{
		ID:        int32(input.ID),
		UpdatedBy: &accId,
		UpdatedAt: &now,

		Name:            input.Name,
		Description:     input.Description,
		Code:            input.Code,
		QueueNumber:     input.QueueNumber,
		IconColor:       input.IconColor,
		IdCustomSvgIcon: input.IdCustomSvgIcon,
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

func (s *service) DeleteTemplRecomendedFrequancyByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteTemplRecomendedFrequancyByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
