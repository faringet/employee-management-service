package endpoints

import (
	"context"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type GetTemplQuestionaryTagsRequest struct {
	PaginationRequest
}

type CreateTemplQuestionaryTagsRequest struct {
	CreatorEmail          string
	TagId                 int32 `json:"tag_id"`
	TemplateQuestionaryId int32 `json:"template_questionary_id"`
}

type CreateTemplQuestionaryTagsRangeRequest struct {
	CreatorEmail         string
	TemplQuestionaryTags []CreateTemplQuestionaryTagsRequest `json:"templ_questionary_tags"`
}

type UpdateTemplQuestionaryTagsRequest struct {
	UpdaterEmail          string
	ID                    int32 `json:"id"`
	TagId                 int32 `json:"tag_id"`
	TemplateQuestionaryId int32 `json:"template_questionary_id"`
}

// Profile godoc
// Get TemplQuestionaryTags
//
//	@Summary		Get TemplQuestionaryTags
//	@Description	Get TemplQuestionaryTags
//	@Tags			TemplQuestionaryTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetTemplQuestionaryTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/survey_tags/{id} [get]
func MakeGetTemplQuestionaryTagsByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BaseIdRequest)
		res, err := s.GetTemplQuestionaryTagsByID(ctx, templateservice.BaseIdRequest{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return res, err
		// return GetTemplQuestionaryTagsResponse{
		// 	ID:                    res.ID,
		// 	UpdatedAt:             res.UpdatedAt,
		// 	CreatedAt:             res.CreatedAt,
		// 	UpdatedBy:             res.UpdatedBy,
		// 	CreatedBy:             res.CreatedBy,
		// 	TagId:                 res.TagId,
		// 	TemplateQuestionaryId: res.TemplateQuestionaryId,
		// 	SurveyTags: &BaseSurveyTagsDTO{
		// 		ID:          res.SurveyTags.ID,
		// 		Name:        res.SurveyTags.Name,
		// 		Code:        res.SurveyTags.Code,
		// 		Description: res.SurveyTags.Description,
		// 	},
		// }, nil
	}
}

// Profile godoc
// Get TemplQuestionaryTags
//
//	@Summary		Get TemplQuestionaryTags
//	@Description	Get TemplQuestionaryTags
//	@Tags			TemplQuestionaryTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetTemplQuestionaryTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [get]
func MakeGetTemplQuestionaryTagsEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTemplQuestionaryTagsRequest)

		input := templateservice.GetTemplQuestionaryTagsRequest{
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

		res, err := s.GetTemplQuestionaryTags(ctx, input)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// Profile godoc
// Create TemplQuestionaryTags
//
//	@Summary		Create TemplQuestionaryTags
//	@Description	Create TemplQuestionaryTags
//	@Tags			TemplQuestionaryTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	CreateTemplQuestionaryTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [post]
func MakeCreateTemplQuestionaryTagsEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTemplQuestionaryTagsRequest)

		input := templateservice.BaseTemplQuestionaryTagsDTO{
			ModifierEmail:         req.CreatorEmail,
			TagId:                 req.TagId,
			TemplateQuestionaryId: req.TemplateQuestionaryId,
		}

		res, err := s.CreateTemplQuestionaryTags(ctx, input)
		if err != nil {
			return nil, err
		}

		return BaseIdRequest(*res), nil
	}
}

// Profile godoc
// Create TemplQuestionaryTags
//
//	@Summary		Create TemplQuestionaryTags
//	@Description	Create TemplQuestionaryTags
//	@Tags			TemplQuestionaryTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	CreateTemplQuestionaryTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [post]
func MakeCreateTemplQuestionaryTagsRangeEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTemplQuestionaryTagsRangeRequest)

		var input []templateservice.BaseTemplQuestionaryTagsDTO
		for _, r := range req.TemplQuestionaryTags {
			input = append(input, templateservice.BaseTemplQuestionaryTagsDTO{
				ModifierEmail:         r.CreatorEmail,
				TagId:                 r.TagId,
				TemplateQuestionaryId: r.TemplateQuestionaryId,
			})
		}

		err := s.CreateTemplQuestionaryTagsRange(ctx, input)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

// Profile godoc
// Update TemplQuestionaryTags
//
//	@Summary		Update TemplQuestionaryTags
//	@Description	Update TemplQuestionaryTags
//	@Tags			TemplQuestionaryTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	UpdateTemplQuestionaryTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries{id} [put]
func MakeUpdateTemplQuestionaryTagsEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTemplQuestionaryTagsRequest)

		input := templateservice.BaseTemplQuestionaryTagsDTO{
			ID:                    req.ID,
			TagId:                 req.TagId,
			TemplateQuestionaryId: req.TemplateQuestionaryId,
		}

		res, err := s.UpdateTemplQuestionaryTagsByID(ctx, input)
		if err != nil {
			return nil, err
		}

		return BaseIdResponse(*res), nil
	}
}

// Profile godoc
// Delete TemplQuestionaryTags
//
//	@Summary		Delete TemplQuestionaryTags
//	@Description	Delete TemplQuestionaryTags
//	@Tags			TemplQuestionaryTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	DeleteTemplQuestionaryTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries{id} [delete]
func MakeDeleteTemplQuestionaryTagsByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BaseIdRequest)

		input := templateservice.BaseIdRequest(req)

		err := s.DeleteTemplQuestionaryTagsByID(ctx, input)
		if err != nil {
			return nil, errors.ErrInternalServerError
		}

		return nil, nil
	}
}
