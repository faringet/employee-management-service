package endpoints

import (
	"context"
	"net/http"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type DeleteCommunicationTemplateByIDRequest struct {
	ID int
}

// Profile godoc
// Delete communication template
//
//	@Summary		Delete communication template
//	@Description	Delete communication template record by id
//	@Tags			communication
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"ID"
//
//	@Success		200	{object}	EmptyResponseSample	"json with empty response"
//	@Failure		400	{object}	SwaggerSimpleError	"json with error msg"
//	@Failure		401	{object}	SwaggerError		"json with error msg"
//	@Failure		403	{object}	SwaggerSimpleError	"json with error msg"
//	@Failure		404	{object}	SwaggerSimpleError	"json with error msg"
//	@Failure		422	{object}	SwaggerError		"json with error msg"
//	@Failure		500	{object}	SwaggerSimpleError	"json with error msg"
//	@Security		Bearer
//
//	@Router			/communication/{id} [delete]
func MakeDeleteCommunicationTemplateByIDEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeleteCommunicationTemplateByIDRequest)
		if !ok {
			return nil, ie.Error{
				Code:    http.StatusBadRequest,
				Message: "request has incorrect structure",
			}
		}

		deleteCommunicationTemplateReq := templateservice.DeleteCommunicationTemplateByIDRequest(req)
		err := s.DeleteCommunicationTemplateByID(ctx, &deleteCommunicationTemplateReq)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
