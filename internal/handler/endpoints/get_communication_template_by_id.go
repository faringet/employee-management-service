package endpoints

import (
	"context"
	"net/http"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type GetCommunicationTemplateByID struct {
	ID int
}

// Profile godoc
// Get communication template
//
//	@Summary		Get communication template
//	@Description	Get communication template record by id
//	@Tags			communication
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"ID"
//
//	@Success		200	{object}	CommunicationTemplateResponse	"json with empty response"
//	@Failure		400	{object}	SwaggerSimpleError				"json with error msg"
//	@Failure		401	{object}	SwaggerError					"json with error msg"
//	@Failure		403	{object}	SwaggerSimpleError				"json with error msg"
//	@Failure		404	{object}	SwaggerSimpleError				"json with error msg"
//	@Failure		422	{object}	SwaggerError					"json with error msg"
//	@Failure		500	{object}	SwaggerSimpleError				"json with error msg"
//	@Security		Bearer
//
//	@Router			/communication/{id} [get]
func MakeGetCommunicationTemplateByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetCommunicationTemplateByID)
		if !ok {
			return nil, ie.Error{
				Code:    http.StatusBadRequest,
				Message: "request has incorrect structure",
			}
		}

		getCommunicationTemplateReq := templateservice.GetCommunicationTemplateByIDRequest(req)
		templateComms, err := s.GetCommunicationTemplateByID(ctx, &getCommunicationTemplateReq)
		if err != nil {
			return nil, err
		}

		resp := CommunicationTemplateResponse(*templateComms)
		return resp, nil
	}
}
