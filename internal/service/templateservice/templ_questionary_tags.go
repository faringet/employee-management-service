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

type TemplQuestionaryTags interface {
	CreateTemplQuestionaryTags(ctx context.Context, input BaseTemplQuestionaryTagsDTO) (*BaseIdResponse, error)
	CreateTemplQuestionaryTagsRange(ctx context.Context, input []BaseTemplQuestionaryTagsDTO) error
	DeleteTemplQuestionaryTagsByID(ctx context.Context, input BaseIdRequest) error
	GetTemplQuestionaryTagsByID(ctx context.Context, req BaseIdRequest) (*BaseTemplQuestionaryTagsDTO, error)
	GetTemplQuestionaryTags(ctx context.Context, req GetTemplQuestionaryTagsRequest) ([]BaseTemplQuestionaryTagsDTO, error)
	UpdateTemplQuestionaryTagsByID(ctx context.Context, input BaseTemplQuestionaryTagsDTO) (*BaseIdResponse, error)
}

type BaseTemplQuestionaryTagsDTO struct {
	ModifierEmail         string
	ID                    int32
	TagId                 int32
	TemplateQuestionaryId int32
	UpdatedAt             *time.Time
	CreatedAt             *time.Time
	CreatedBy             *int32
	UpdatedBy             *int32
	SurveyTags            *BaseSurveyTagsDTO
}

type GetTemplQuestionaryTagsRequest struct {
	PaginationRequest
}

type CreateTemplQuestionaryTagsRangeDTO struct {
	CreatorEmail         string
	TemplQuestionaryTags []BaseTemplQuestionaryTagsDTO
}

func (s *service) CreateTemplQuestionaryTags(ctx context.Context, input BaseTemplQuestionaryTagsDTO) (*BaseIdResponse, error) {
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

	req := model.TemplQuestionaryTags{
		CreatedBy:             &accId,
		UpdatedBy:             &accId,
		TagID:                 int32(input.TagId),
		TemplateQuestionaryID: &input.TemplateQuestionaryId,
	}

	res, err := s.repo.CreateTemplQuestionaryTags(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateTemplQuestionaryTagsRange(ctx context.Context, input []BaseTemplQuestionaryTagsDTO) error {
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
	var rows []model.TemplQuestionaryTags
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.TemplQuestionaryTags{
			CreatedBy:             &accId,
			UpdatedBy:             &accId,
			TagID:                 row.TagId,
			TemplateQuestionaryID: &row.TemplateQuestionaryId,
		})
	}

	err := s.repo.CreateTemplQuestionaryTagsRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetTemplQuestionaryTagsByID(ctx context.Context, req BaseIdRequest) (*BaseTemplQuestionaryTagsDTO, error) {
	row, err := s.repo.GetTemplQuestionaryTagsByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseTemplQuestionaryTagsDTO{
		ID:                    row.ID,
		UpdatedAt:             row.UpdatedAt,
		CreatedAt:             row.CreatedAt,
		UpdatedBy:             row.UpdatedBy,
		CreatedBy:             row.CreatedBy,
		TagId:                 row.TagID,
		TemplateQuestionaryId: *row.TemplateQuestionaryID,
		SurveyTags: &BaseSurveyTagsDTO{
			ID:   row.SurveyTags.ID,
			Name: row.SurveyTags.Name,
			Code: row.SurveyTags.Code,
		},
	}

	return tq, nil
}

func (s *service) GetTemplQuestionaryTags(ctx context.Context, req GetTemplQuestionaryTagsRequest) ([]BaseTemplQuestionaryTagsDTO, error) {

	input := repomodel.GetTemplQuestionaryTagsRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetTemplQuestionaryTags(ctx, input)
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

	res := []BaseTemplQuestionaryTagsDTO{}
	for _, row := range rows {
		tq := BaseTemplQuestionaryTagsDTO{
			ID:                    row.ID,
			UpdatedAt:             row.UpdatedAt,
			CreatedAt:             row.CreatedAt,
			UpdatedBy:             row.UpdatedBy,
			CreatedBy:             row.CreatedBy,
			TagId:                 row.TagID,
			TemplateQuestionaryId: *row.TemplateQuestionaryID,

			SurveyTags: &BaseSurveyTagsDTO{
				ID:   row.SurveyTags.ID,
				Name: row.SurveyTags.Name,
				Code: row.SurveyTags.Code,
			},
		}

		res = append(res, tq)
	}

	return res, nil
}

func (s *service) UpdateTemplQuestionaryTagsByID(ctx context.Context, input BaseTemplQuestionaryTagsDTO) (*BaseIdResponse, error) {
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

	res, err := s.repo.UpdateTemplQuestionaryTagsByID(ctx, model.TemplQuestionaryTags{
		ID:                    input.ID,
		UpdatedBy:             &accId,
		TagID:                 input.TagId,
		TemplateQuestionaryID: &input.TemplateQuestionaryId,
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

func (s *service) DeleteTemplQuestionaryTagsByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteTemplQuestionaryTagsByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
