package endpoints

import (
	"context"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type GetTemplRecomendedFrequancyRequest struct {
	PaginationRequest
}

type CreateTemplRecomendedFrequancyRequest struct {
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

type CreateTemplRecomendedFrequancyRangeRequest struct {
	CreatorEmail             string
	TemplRecomendedFrequancy []CreateTemplRecomendedFrequancyRequest `json:"survey_tags"`
}

type UpdateTemplRecomendedFrequancyRequest struct {
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
// Get TemplRecomendedFrequancy
//
//	@Summary		Get TemplRecomendedFrequancy
//	@Description	Get TemplRecomendedFrequancy
//	@Tags			TemplRecomendedFrequancy
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetTemplRecomendedFrequancyResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/survey_tags/{id} [get]
func MakeGetTemplRecomendedFrequancyByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BaseIdRequest)
		res, err := s.GetTemplRecomendedFrequancyByID(ctx, templateservice.BaseIdRequest{
			ID: req.ID,
		})
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// Profile godoc
// Get TemplRecomendedFrequancy
//
//	@Summary		Get TemplRecomendedFrequancy
//	@Description	Get TemplRecomendedFrequancy
//	@Tags			TemplRecomendedFrequancy
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	GetTemplRecomendedFrequancyResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [get]
func MakeGetTemplRecomendedFrequancyEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTemplRecomendedFrequancyRequest)

		input := templateservice.GetTemplRecomendedFrequancyRequest{
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

		res, err := s.GetTemplRecomendedFrequancy(ctx, input)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// Profile godoc
// Create TemplRecomendedFrequancy
//
//	@Summary		Create TemplRecomendedFrequancy
//	@Description	Create TemplRecomendedFrequancy
//	@Tags			TemplRecomendedFrequancy
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	CreateTemplRecomendedFrequancyResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [post]
func MakeCreateTemplRecomendedFrequancyEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTemplRecomendedFrequancyRequest)

		input := templateservice.BaseTemplRecomendedFrequancyDTO{
			Name:            req.Name,
			Code:            req.Code,
			Description:     req.Description,
			QueueNumber:     req.QueueNumber,
			IconColor:       req.IconColor,
			IdCustomSvgIcon: req.IdCustomSvgIcon,
			CreatedBy:       req.CreatedBy,
			UpdatedBy:       req.UpdatedBy,
		}

		res, err := s.CreateTemplRecomendedFrequancy(ctx, input)
		if err != nil {
			return nil, err
		}

		return BaseIdResponse(*res), nil
	}
}

// Profile godoc
// Create TemplRecomendedFrequancy
//
//	@Summary		Create TemplRecomendedFrequancy
//	@Description	Create TemplRecomendedFrequancy
//	@Tags			TemplRecomendedFrequancy
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	CreateTemplRecomendedFrequancyResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries [post]
func MakeCreateTemplRecomendedFrequancyRangeEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTemplRecomendedFrequancyRangeRequest)

		var input []templateservice.BaseTemplRecomendedFrequancyDTO
		for _, req := range req.TemplRecomendedFrequancy {
			input = append(input, templateservice.BaseTemplRecomendedFrequancyDTO{
				Name:            req.Name,
				Code:            req.Code,
				Description:     req.Description,
				QueueNumber:     req.QueueNumber,
				IconColor:       req.IconColor,
				IdCustomSvgIcon: req.IdCustomSvgIcon,
				UpdatedBy:       req.UpdatedBy,
			})
		}

		err := s.CreateTemplRecomendedFrequancyRange(ctx, input)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

// Profile godoc
// Update TemplRecomendedFrequancy
//
//	@Summary		Update TemplRecomendedFrequancy
//	@Description	Update TemplRecomendedFrequancy
//	@Tags			TemplRecomendedFrequancy
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	UpdateTemplRecomendedFrequancyResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries{id} [put]
func MakeUpdateTemplRecomendedFrequancyEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTemplRecomendedFrequancyRequest)

		input := templateservice.BaseTemplRecomendedFrequancyDTO{
			ID:              req.ID,
			Name:            req.Name,
			Code:            req.Code,
			Description:     req.Description,
			QueueNumber:     req.QueueNumber,
			IconColor:       req.IconColor,
			IdCustomSvgIcon: req.IdCustomSvgIcon,
			UpdatedBy:       req.UpdatedBy,
		}

		res, err := s.UpdateTemplRecomendedFrequancyByID(ctx, input)
		if err != nil {
			return nil, err
		}

		return BaseIdResponse(*res), nil
	}
}

// Profile godoc
// Delete TemplRecomendedFrequancy
//
//	@Summary		Delete TemplRecomendedFrequancy
//	@Description	Delete TemplRecomendedFrequancy
//	@Tags			TemplRecomendedFrequancy
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	DeleteTemplRecomendedFrequancyResponse
//	@Failure		400	{object}	SwaggerSimpleError
//	@Failure		403	{object}	SwaggerSimpleError
//	@Failure		404	{object}	SwaggerSimpleError
//	@Failure		422	{object}	SwaggerError
//	@Failure		500	{object}	SwaggerSimpleError
//	@Security		Bearer
//
//	@Router			/template/questionaries{id} [delete]
func MakeDeleteTemplRecomendedFrequancyByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BaseIdRequest)

		input := templateservice.BaseIdRequest(req)

		err := s.DeleteTemplRecomendedFrequancyByID(ctx, input)
		if err != nil {
			return nil, errors.ErrInternalServerError
		}

		return nil, nil
	}
}
