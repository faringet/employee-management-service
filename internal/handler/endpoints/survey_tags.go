package endpoints

import (
	"context"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type GetSurveyTagsRequest struct {
	PaginationRequest
}

type CreateSurveyTagsRequest struct {
	CreatorEmail    string
	Name            *string `json:"name"`
	Code            *string `json:"code"`
	Description     *string `json:"description"`
	QueueNumber     *int32  `json:"queueNumber"`
	IconColor       *string `json:"iconColor"`
	IdCustomSvgIcon *int32  `json:"idCustomSvgIcon"`
	CreatedBy       *int32  `json:"created_by"`
	UpdatedBy       *int32  `json:"updated_by"`
}

type CreateSurveyTagsRangeRequest struct {
	CreatorEmail string
	SurveyTags   []CreateSurveyTagsRequest `json:"survey_tags"`
}

type UpdateSurveyTagsRequest struct {
	UpdaterEmail    string
	ID              int32   `json:"id"`
	Name            *string `json:"name"`
	Code            *string `json:"code"`
	Description     *string `json:"description"`
	QueueNumber     *int32  `json:"queueNumber"`
	IconColor       *string `json:"iconColor"`
	IdCustomSvgIcon *int32  `json:"idCustomSvgIcon"`
	UpdatedBy       *int32  `json:"updated_by"`
}

// Profile godoc
// Get SurveyTags
//
//	@Summary		Get SurveyTags
//	@Description	Get SurveyTags
//	@Tags			SurveyTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetSurveyTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/survey_tags/{id} [get]
func MakeGetSurveyTagsByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BaseIdRequest)
		res, err := s.GetSurveyTagsByID(ctx, templateservice.BaseIdRequest{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// Profile godoc
// Get SurveyTags
//
//	@Summary		Get SurveyTags
//	@Description	Get SurveyTags
//	@Tags			SurveyTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetSurveyTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [get]
func MakeGetSurveyTagsEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := s.GetSurveyTags(ctx)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// Profile godoc
// Get SurveyTags
//
//	@Summary		Get SurveyTags
//	@Description	Get SurveyTags
//	@Tags			SurveyTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetSurveyTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [get]
func MakeGetSurveyTagsPaginationEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSurveyTagsRequest)

		input := templateservice.GetSurveyTagsRequest{
			PaginationRequest: templateservice.PaginationRequest{
				SearchBy:           []string{req.SearchBy}, //by default we use only one field, but you can modify for multiple
				SearchValue:        []string{req.SearchValue},
				SortBy:             req.SortBy,
				SortType:           req.SortType,
				SearchLogicOpeator: req.SearchLogicOpeator,
				Limit:              req.Limit,
				Offset:             req.Offset,
			},
		}

		res, err := s.GetSurveyTagsPagination(ctx, input)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// Profile godoc
// Create SurveyTags
//
//	@Summary		Create SurveyTags
//	@Description	Create SurveyTags
//	@Tags			SurveyTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	CreateSurveyTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [post]
func MakeCreateSurveyTagsEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSurveyTagsRequest)

		input := templateservice.BaseSurveyTagsDTO{
			Name:            req.Name,
			Code:            req.Code,
			Description:     req.Description,
			QueueNumber:     req.QueueNumber,
			IconColor:       req.IconColor,
			IdCustomSvgIcon: req.IdCustomSvgIcon,
			CreatedBy:       req.CreatedBy,
			UpdatedBy:       req.UpdatedBy,
		}

		res, err := s.CreateSurveyTags(ctx, input)
		if err != nil {
			return nil, err
		}

		return BaseIdResponse(*res), nil
	}
}

// Profile godoc
// Create SurveyTags
//
//	@Summary		Create SurveyTags
//	@Description	Create SurveyTags
//	@Tags			SurveyTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	CreateSurveyTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [post]
func MakeCreateSurveyTagsRangeEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSurveyTagsRangeRequest)

		var input []templateservice.BaseSurveyTagsDTO
		for _, req := range req.SurveyTags {
			input = append(input, templateservice.BaseSurveyTagsDTO{
				Name:            req.Name,
				Code:            req.Code,
				Description:     req.Description,
				QueueNumber:     req.QueueNumber,
				IconColor:       req.IconColor,
				IdCustomSvgIcon: req.IdCustomSvgIcon,
				UpdatedBy:       req.UpdatedBy,
			})
		}

		err := s.CreateSurveyTagsRange(ctx, input)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

// Profile godoc
// Update SurveyTags
//
//	@Summary		Update SurveyTags
//	@Description	Update SurveyTags
//	@Tags			SurveyTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	UpdateSurveyTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries{id} [put]
func MakeUpdateSurveyTagsEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateSurveyTagsRequest)

		input := templateservice.BaseSurveyTagsDTO{
			ID:              req.ID,
			Name:            req.Name,
			Code:            req.Code,
			Description:     req.Description,
			QueueNumber:     req.QueueNumber,
			IconColor:       req.IconColor,
			IdCustomSvgIcon: req.IdCustomSvgIcon,
			UpdatedBy:       req.UpdatedBy,
		}

		res, err := s.UpdateSurveyTagsByID(ctx, input)
		if err != nil {
			return nil, err
		}

		return BaseIdResponse(*res), nil
	}
}

// Profile godoc
// Delete SurveyTags
//
//	@Summary		Delete SurveyTags
//	@Description	Delete SurveyTags
//	@Tags			SurveyTags
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	DeleteSurveyTagsResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries{id} [delete]
func MakeDeleteSurveyTagsByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BaseIdRequest)

		input := templateservice.BaseIdRequest(req)

		err := s.DeleteSurveyTagsByID(ctx, input)
		if err != nil {
			return nil, errors.ErrInternalServerError
		}

		return nil, nil
	}
}
