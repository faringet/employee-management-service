package endpoints

import (
	"context"
	"net/http"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type GetCommunicationTemplatesByEntityID struct {
	ID int
}

// Profile godoc
// Get communication templates by entity id
//
//	@Summary		Get communication templates
//	@Description	Get communication template records by entity id
//	@Tags			communication
//	@Produce		json
//	@Param			id	path		int								true	"ID"
//
//	@Success		200	{array}		CommunicationTemplateResponse	"json with empty response"
//	@Failure		400	{object}	SwaggerSimpleError				"json with error msg"
//	@Failure		401	{object}	SwaggerError					"json with error msg"
//	@Failure		403	{object}	SwaggerSimpleError				"json with error msg"
//	@Failure		404	{object}	SwaggerSimpleError				"json with error msg"
//	@Failure		422	{object}	SwaggerError					"json with error msg"
//	@Failure		500	{object}	SwaggerSimpleError				"json with error msg"
//	@Security		Bearer
//
//	@Router			/communication/entity/{id} [get]
func MakeGetCommunicationTemplatesByEntityIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetCommunicationTemplatesByEntityID)
		if !ok {
			return nil, ie.Error{
				Code:    http.StatusBadRequest,
				Message: "request has incorrect structure",
			}
		}

		communicationTemplatesReq := templateservice.GetCommunicationTemplatesByEntityIDRequest(req)
		resp, err := s.GetCommunicationTemplatesByEntityID(ctx, &communicationTemplatesReq)
		if err != nil {
			return nil, err
		}

		var res []CommunicationTemplateResponse
		for _, template := range resp {
			res = append(res, CommunicationTemplateResponse(template))
		}
		return res, nil
	}
}
