package endpoints

import (
	"context"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type CreateTemplTemplateQuestionaryRequest struct {
	CreatorEmail string

	IsTemplateER          *bool   `json:"istemplateER"`
	Name                  *string `json:"name"`
	Estimation            *string `json:"estimation"`
	Description           *string `json:"descriptiopn"`
	Survey                *string `json:"survey"`
	RecomendedFrequencyId *int32  `json:"recomended_frequency_id"`
}

type GetTemplTemplateQuestionaryRequest struct {
	PaginationRequest
}

type CreateTemplTemplateQuestionaryRangeRequest struct {
	CreatorEmail             string
	TemplTemplateQuestionary []CreateTemplTemplateQuestionaryRequest `json:"templ_questionary_tags"`
}

type UpdateTemplTemplateQuestionaryRequest struct {
	UpdaterEmail string
	ID           int32 `json:"id"`

	IsTemplateER          *bool   `json:"istemplateER"`
	Name                  *string `json:"name"`
	Estimation            *string `json:"estimation"`
	Description           *string `json:"descriptiopn"`
	Survey                *string `json:"survey"`
	RecomendedFrequencyId *int32  `json:"recomended_frequency_id"`
}

// Profile godoc
// Get TemplTemplateQuestionary
//
//	@Summary		Get TemplTemplateQuestionary
//	@Description	Get TemplTemplateQuestionary
//	@Tags			TemplTemplateQuestionary
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetTemplTemplateQuestionaryResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/survey_tags/{id} [get]
func MakeGetTemplTemplateQuestionaryByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BaseIdRequest)
		res, err := s.GetTemplTemplateQuestionaryByID(ctx, templateservice.BaseIdRequest{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return res, nil
		// return GetTemplTemplateQuestionaryResponse{
		// 	ID:                    res.ID,
		// 	UpdatedAt:             res.UpdatedAt,
		// 	CreatedAt:             res.CreatedAt,
		// 	UpdatedBy:             res.UpdatedBy,
		// 	CreatedBy:             res.CreatedBy,
		// 	TagId:                 res.TagId,
		// 	TemplateQuestionaryId: res.TemplateQuestionaryId,
		// 	SurveyTags: &GetSurveyTagsResponse{
		// 		ID:          res.SurveyTags.ID,
		// 		Name:        res.SurveyTags.Name,
		// 		Code:        res.SurveyTags.Code,
		// 		Description: res.SurveyTags.Description,
		// 	},
		// }, nil
	}
}

// Profile godoc
// Get TemplTemplateQuestionary
//
//	@Summary		Get TemplTemplateQuestionary
//	@Description	Get TemplTemplateQuestionary
//	@Tags			TemplTemplateQuestionary
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetTemplTemplateQuestionaryResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [get]
func MakeGetTemplTemplateQuestionaryEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetTemplTemplateQuestionaryRequest)

		input := templateservice.GetTemplTemplateQuestionaryRequest{
			PaginationRequest: templateservice.PaginationRequest{
				SearchBy:           []string{req.SearchBy},
				SearchValue:        []string{req.SearchValue},
				SortBy:             req.SortBy,
				SortType:           req.SortType,
				SearchLogicOpeator: req.SearchLogicOpeator,
				Limit:              req.Limit,
				Offset:             req.Offset,
			},
		}

		res, err := s.GetTemplTemplateQuestionary(ctx, input)
		if err != nil {
			return nil, err
		}

		return res, err
		// tqs := []GetTemplTemplateQuestionaryResponse{}
		// for _, r := range res {
		// 	tqs = append(tqs, GetTemplTemplateQuestionaryResponse{
		// 		ID:                    r.ID,
		// 		UpdatedAt:             r.UpdatedAt,
		// 		CreatedAt:             r.CreatedAt,
		// 		UpdatedBy:             r.UpdatedBy,
		// 		CreatedBy:             r.CreatedBy,
		// 		TagId:                 r.TagId,
		// 		TemplateQuestionaryId: r.TemplateQuestionaryId,
		// 		SurveyTags: &GetSurveyTagsResponse{
		// 			ID:          r.SurveyTags.ID,
		// 			Name:        r.SurveyTags.Name,
		// 			Code:        r.SurveyTags.Code,
		// 			Description: r.SurveyTags.Description,
		// 		},
		// 	})
		// }
		// return tqs, nil
	}
}

// Profile godoc
// Create TemplTemplateQuestionary
//
//	@Summary		Create TemplTemplateQuestionary
//	@Description	Create TemplTemplateQuestionary
//	@Tags			TemplTemplateQuestionary
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	CreateTemplTemplateQuestionaryResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [post]
func MakeCreateTemplTemplateQuestionaryEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTemplTemplateQuestionaryRequest)

		input := templateservice.BaseTemplTemplateQuestionaryDTO{
			ModifierEmail: req.CreatorEmail,

			Name:                  req.Name,
			Description:           req.Description,
			Survey:                req.Survey,
			Estimation:            req.Estimation,
			IsTemplateER:          req.IsTemplateER,
			RecomendedFrequencyId: req.RecomendedFrequencyId,
		}

		res, err := s.CreateTemplTemplateQuestionary(ctx, input)
		if err != nil {
			return nil, err
		}

		return BaseIdRequest(*res), nil
	}
}

// Profile godoc
// Create TemplTemplateQuestionary
//
//	@Summary		Create TemplTemplateQuestionary
//	@Description	Create TemplTemplateQuestionary
//	@Tags			TemplTemplateQuestionary
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	CreateTemplTemplateQuestionaryResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [post]
func MakeCreateTemplTemplateQuestionaryRangeEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTemplTemplateQuestionaryRangeRequest)

		var input []templateservice.BaseTemplTemplateQuestionaryDTO
		for _, r := range req.TemplTemplateQuestionary {
			input = append(input, templateservice.BaseTemplTemplateQuestionaryDTO{
				ModifierEmail: r.CreatorEmail,

				Name:                  r.Name,
				Description:           r.Description,
				Survey:                r.Survey,
				Estimation:            r.Estimation,
				IsTemplateER:          r.IsTemplateER,
				RecomendedFrequencyId: r.RecomendedFrequencyId,
			})
		}

		err := s.CreateTemplTemplateQuestionaryRange(ctx, input)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

// Profile godoc
// Update TemplTemplateQuestionary
//
//	@Summary		Update TemplTemplateQuestionary
//	@Description	Update TemplTemplateQuestionary
//	@Tags			TemplTemplateQuestionary
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	UpdateTemplTemplateQuestionaryResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries{id} [put]
func MakeUpdateTemplTemplateQuestionaryEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTemplTemplateQuestionaryRequest)

		input := templateservice.BaseTemplTemplateQuestionaryDTO{
			ID: req.ID,

			Name:                  req.Name,
			Description:           req.Description,
			Survey:                req.Survey,
			Estimation:            req.Estimation,
			IsTemplateER:          req.IsTemplateER,
			RecomendedFrequencyId: req.RecomendedFrequencyId,
		}

		res, err := s.UpdateTemplTemplateQuestionaryByID(ctx, input)
		if err != nil {
			return nil, err
		}

		return BaseIdResponse(*res), nil
	}
}

// Profile godoc
// Delete TemplTemplateQuestionary
//
//	@Summary		Delete TemplTemplateQuestionary
//	@Description	Delete TemplTemplateQuestionary
//	@Tags			TemplTemplateQuestionary
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	DeleteTemplTemplateQuestionaryResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries{id} [delete]
func MakeDeleteTemplTemplateQuestionaryByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BaseIdRequest)

		input := templateservice.BaseIdRequest(req)

		err := s.DeleteTemplTemplateQuestionaryByID(ctx, input)
		if err != nil {
			return nil, errors.ErrInternalServerError
		}

		return nil, nil
	}
}
