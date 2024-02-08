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

type TemplTemplateQuestionary interface {
	CreateTemplTemplateQuestionary(ctx context.Context, input BaseTemplTemplateQuestionaryDTO) (*BaseIdResponse, error)
	CreateTemplTemplateQuestionaryRange(ctx context.Context, input []BaseTemplTemplateQuestionaryDTO) error
	DeleteTemplTemplateQuestionaryByID(ctx context.Context, input BaseIdRequest) error
	GetTemplTemplateQuestionaryByID(ctx context.Context, req BaseIdRequest) (*BaseTemplTemplateQuestionaryDTO, error)
	GetTemplTemplateQuestionary(ctx context.Context, req GetTemplTemplateQuestionaryRequest) ([]BaseTemplTemplateQuestionaryDTO, error)
	UpdateTemplTemplateQuestionaryByID(ctx context.Context, input BaseTemplTemplateQuestionaryDTO) (*BaseIdResponse, error)
}

type BaseTemplTemplateQuestionaryDTO struct {
	ModifierEmail string
	ID            int32      `json:"id"`
	UpdatedAt     *time.Time `json:"updated_at"` //как в БД
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedBy     *int32     `json:"updated_by"`
	CreatedBy     *int32     `json:"created_by"`

	Name                  *string `json:"name"` //как в БД
	Description           *string `json:"description"`
	Estimation            *string `json:"estimation"`
	IsTemplateER          *bool   `json:"istemplateER"`
	Survey                *string `json:"survey"`
	RecomendedFrequencyId *int32  `json:"recomended_frequency_id"` //как в БД

	TemplRecomendedFrequancy BaseTemplRecomendedFrequancyDTO `json:"templ_recomended_frequancy"` //как в БД
	TemplQuestionaryTags     []BaseTemplQuestionaryTagsDTO   `json:"templ_questionary_tags"`     //как в БД
}

type GetTemplTemplateQuestionaryRequest struct {
	PaginationRequest
}

type CreateTemplTemplateQuestionaryRangeDTO struct {
	CreatorEmail             string
	TemplTemplateQuestionary []BaseTemplTemplateQuestionaryDTO
}

func (s *service) CreateTemplTemplateQuestionary(ctx context.Context, input BaseTemplTemplateQuestionaryDTO) (*BaseIdResponse, error) {
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

	req := model.TemplTemplateQuestionary{
		CreatedBy:             &accId,
		UpdatedBy:             &accId,
		Description:           input.Description,
		Name:                  input.Name,
		Estimation:            input.Estimation,
		IsTemplateER:          input.IsTemplateER,
		Survey:                input.Survey,
		RecomendedFrequencyID: input.RecomendedFrequencyId,
	}

	res, err := s.repo.CreateTemplTemplateQuestionary(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateTemplTemplateQuestionaryRange(ctx context.Context, input []BaseTemplTemplateQuestionaryDTO) error {
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
	var rows []model.TemplTemplateQuestionary
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.TemplTemplateQuestionary{
			CreatedBy:             &accId,
			UpdatedBy:             &accId,
			Description:           row.Description,
			Name:                  row.Name,
			Estimation:            row.Estimation,
			IsTemplateER:          row.IsTemplateER,
			Survey:                row.Survey,
			RecomendedFrequencyID: row.RecomendedFrequencyId,
		})
	}

	err := s.repo.CreateTemplTemplateQuestionaryRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetTemplTemplateQuestionaryByID(ctx context.Context, req BaseIdRequest) (*BaseTemplTemplateQuestionaryDTO, error) {
	row, err := s.repo.GetTemplTemplateQuestionaryByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseTemplTemplateQuestionaryDTO{
		ID:           row.ID,
		UpdatedAt:    row.UpdatedAt,
		CreatedAt:    row.CreatedAt,
		UpdatedBy:    row.UpdatedBy,
		CreatedBy:    row.CreatedBy,
		Name:         row.Name,
		Description:  row.Description,
		Estimation:   row.Estimation,
		IsTemplateER: row.IsTemplateER,

		TemplRecomendedFrequancy: BaseTemplRecomendedFrequancyDTO{
			ID:              row.TemplRecomendedFrequancy.ID,
			Name:            row.TemplRecomendedFrequancy.Name,
			Code:            row.TemplRecomendedFrequancy.Code,
			Description:     row.TemplRecomendedFrequancy.Description,
			QueueNumber:     row.TemplRecomendedFrequancy.QueueNumber,
			IdCustomSvgIcon: row.TemplRecomendedFrequancy.IdCustomSvgIcon,
			CreatedAt:       row.TemplRecomendedFrequancy.CreatedAt,
			CreatedBy:       row.TemplRecomendedFrequancy.CreatedBy,
			UpdatedAt:       row.TemplRecomendedFrequancy.UpdatedAt,
			UpdatedBy:       row.TemplRecomendedFrequancy.UpdatedBy,
		},
		// Survey:       row.Survey,
	}

	questionaryTags := []BaseTemplQuestionaryTagsDTO{}
	for _, r := range row.TemplQuestionaryTags {
		questionaryTags = append(questionaryTags, BaseTemplQuestionaryTagsDTO{
			ID:        r.ID,
			CreatedAt: r.CreatedAt,
			CreatedBy: r.CreatedBy,
			UpdatedAt: r.UpdatedAt,
			UpdatedBy: r.UpdatedBy,

			TagId:                 r.TagID,
			TemplateQuestionaryId: *r.TemplateQuestionaryID,

			//uncomment if you need 3rd level
			SurveyTags: &BaseSurveyTagsDTO{
				ID:          r.SurveyTags.ID,
				Name:        r.SurveyTags.Name,
				Code:        r.SurveyTags.Code,
				Description: r.SurveyTags.Description,
			},
		})

	}

	tq.TemplQuestionaryTags = questionaryTags

	return tq, nil
}

func (s *service) GetTemplTemplateQuestionary(ctx context.Context, req GetTemplTemplateQuestionaryRequest) ([]BaseTemplTemplateQuestionaryDTO, error) {
	input := repomodel.GetTemplTemplateQuestionaryRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetTemplTemplateQuestionary(ctx, input)
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

	res := []BaseTemplTemplateQuestionaryDTO{}
	for _, row := range rows {
		tq := BaseTemplTemplateQuestionaryDTO{
			ID:        row.ID,
			UpdatedAt: row.UpdatedAt,
			CreatedAt: row.CreatedAt,
			UpdatedBy: row.UpdatedBy,
			CreatedBy: row.CreatedBy,

			Name:         row.Name,
			Description:  row.Description,
			Estimation:   row.Estimation,
			IsTemplateER: row.IsTemplateER,

			TemplRecomendedFrequancy: BaseTemplRecomendedFrequancyDTO{
				ID:              row.TemplRecomendedFrequancy.ID,
				Name:            row.TemplRecomendedFrequancy.Name,
				Code:            row.TemplRecomendedFrequancy.Code,
				Description:     row.TemplRecomendedFrequancy.Description,
				QueueNumber:     row.TemplRecomendedFrequancy.QueueNumber,
				IdCustomSvgIcon: row.TemplRecomendedFrequancy.IdCustomSvgIcon,
				CreatedAt:       row.TemplRecomendedFrequancy.CreatedAt,
				CreatedBy:       row.TemplRecomendedFrequancy.CreatedBy,
				UpdatedAt:       row.TemplRecomendedFrequancy.UpdatedAt,
				UpdatedBy:       row.TemplRecomendedFrequancy.UpdatedBy,
			},

			// Survey:       row.Survey,
		}
		log.Println(row.TemplQuestionaryTags)

		questionaryTags := []BaseTemplQuestionaryTagsDTO{}
		for _, r := range row.TemplQuestionaryTags {

			questionaryTags = append(questionaryTags, BaseTemplQuestionaryTagsDTO{
				ID:        r.ID,
				CreatedAt: r.CreatedAt,
				CreatedBy: r.CreatedBy,
				UpdatedAt: r.UpdatedAt,
				UpdatedBy: r.UpdatedBy,

				TagId:                 r.TagID,
				TemplateQuestionaryId: *r.TemplateQuestionaryID,

				//uncomment if you need 3rd level
				SurveyTags: &BaseSurveyTagsDTO{
					ID:          r.SurveyTags.ID,
					Name:        r.SurveyTags.Name,
					Code:        r.SurveyTags.Code,
					Description: r.SurveyTags.Description,
				},
			})

		}

		tq.TemplQuestionaryTags = questionaryTags
		res = append(res, tq)
	}

	return res, nil
}

func (s *service) UpdateTemplTemplateQuestionaryByID(ctx context.Context, input BaseTemplTemplateQuestionaryDTO) (*BaseIdResponse, error) {
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

	res, err := s.repo.UpdateTemplTemplateQuestionaryByID(ctx, model.TemplTemplateQuestionary{
		ID:        int32(input.ID),
		UpdatedBy: &accId,

		Name:                  input.Name,
		Description:           input.Description,
		Estimation:            input.Estimation,
		IsTemplateER:          input.IsTemplateER,
		Survey:                input.Survey,
		RecomendedFrequencyID: input.RecomendedFrequencyId,
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

func (s *service) DeleteTemplTemplateQuestionaryByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteTemplTemplateQuestionaryByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
