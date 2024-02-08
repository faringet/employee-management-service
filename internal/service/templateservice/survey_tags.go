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

type SurveyTags interface {
	CreateSurveyTags(ctx context.Context, input BaseSurveyTagsDTO) (*BaseIdResponse, error)
	CreateSurveyTagsRange(ctx context.Context, input []BaseSurveyTagsDTO) error
	DeleteSurveyTagsByID(ctx context.Context, input BaseIdRequest) error
	GetSurveyTagsByID(ctx context.Context, req BaseIdRequest) (*BaseSurveyTagsDTO, error)
	GetSurveyTags(ctx context.Context) ([]BaseSurveyTagsDTO, error)
	GetSurveyTagsPagination(ctx context.Context, req GetSurveyTagsRequest) (*GetSurveyTagsResponse, error)
	UpdateSurveyTagsByID(ctx context.Context, input BaseSurveyTagsDTO) (*BaseIdResponse, error)
}

type BaseSurveyTagsDTO struct {
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

type GetSurveyTagsRequest struct {
	PaginationRequest
}

type GetSurveyTagsResponse struct {
	PaginationResponse
	Data []BaseSurveyTagsDTO `json:"data"`
}

type CreateSurveyTagsRangeDTO struct {
	CreatorEmail string
	SurveyTags   []BaseSurveyTagsDTO
}

func (s *service) CreateSurveyTags(ctx context.Context, input BaseSurveyTagsDTO) (*BaseIdResponse, error) {
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

	req := model.SurveyTags{
		CreatedBy:       &accId,
		UpdatedBy:       &accId,
		Name:            input.Name,
		Description:     input.Description,
		Code:            input.Code,
		QueueNumber:     input.QueueNumber,
		IconColor:       input.IconColor,
		IdCustomSvgIcon: input.IdCustomSvgIcon,
	}

	res, err := s.repo.CreateSurveyTags(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateSurveyTagsRange(ctx context.Context, input []BaseSurveyTagsDTO) error {
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
	var rows []model.SurveyTags
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.SurveyTags{
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

	err := s.repo.CreateSurveyTagsRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetSurveyTagsByID(ctx context.Context, req BaseIdRequest) (*BaseSurveyTagsDTO, error) {
	row, err := s.repo.GetSurveyTagsByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseSurveyTagsDTO{
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

func (s *service) GetSurveyTags(ctx context.Context) ([]BaseSurveyTagsDTO, error) {

	rows, err := s.repo.GetSurveyTags(ctx)
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

	res := []BaseSurveyTagsDTO{}
	for _, row := range rows {
		tq := BaseSurveyTagsDTO{
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

func (s *service) GetSurveyTagsPagination(ctx context.Context, req GetSurveyTagsRequest) (*GetSurveyTagsResponse, error) {

	input := repomodel.GetSurveyTagsRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetSurveyTagsPagination(ctx, input)
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

	data := []BaseSurveyTagsDTO{}
	for _, row := range rows {
		tq := BaseSurveyTagsDTO{
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

		data = append(data, tq)
	}

	count, err := s.repo.CountSurveyTags(ctx)
	if err != nil {
		s.logger.Error("count survey_tags", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	res := &GetSurveyTagsResponse{
		Data: data,
		PaginationResponse: PaginationResponse{
			ActivePage:   activePage,
			TotalCount:   *count,
			CountPerPage: req.Limit,
			TotalPages:   totalPages,
		},
	}

	return res, nil
}

func (s *service) UpdateSurveyTagsByID(ctx context.Context, input BaseSurveyTagsDTO) (*BaseIdResponse, error) {
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

	res, err := s.repo.UpdateSurveyTagsByID(ctx, model.SurveyTags{
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

func (s *service) DeleteSurveyTagsByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteSurveyTagsByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
