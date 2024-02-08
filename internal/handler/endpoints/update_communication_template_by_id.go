package endpoints

import (
	"context"
	"net/http"
	"time"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type UpdateCommunicationTemplateByIDRequest struct {
	ID             int        `json:"id" validate:"required"`
	OwnerEntityID  int        `json:"owner_entity_id" validate:"required"`
	HeaderLogoID   int        `json:"header_logo_id" validate:"required"`
	ReminderDaysID int        `json:"reminder_days_id" validate:"required"`
	IsSendReport   *bool      `json:"is_send_report" validate:"required"`
	Name           *string    `json:"name" validate:"required"`
	Description    *string    `json:"description" validate:"required"`
	TimeSendReport *time.Time `json:"time_send_report" validate:"required"`
	UpdatedBy      *int       `json:"updated_by" validate:"required"`
}

// Profile godoc
// Update communication template
//
//	@Summary		Update communication template
//	@Description	Update communication template record by id
//	@Tags			communication
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int										true	"ID"
//	@Param			request	body		UpdateCommunicationTemplateByIDRequest	true	"Update request"
//
//	@Success		200		{object}	CommunicationTemplateResponse			"json with empty response"
//	@Failure		400		{object}	SwaggerSimpleError						"json with error msg"
//	@Failure		401		{object}	SwaggerError							"json with error msg"
//	@Failure		403		{object}	SwaggerSimpleError						"json with error msg"
//	@Failure		404		{object}	SwaggerSimpleError						"json with error msg"
//	@Failure		422		{object}	SwaggerError							"json with error msg"
//	@Failure		500		{object}	SwaggerSimpleError						"json with error msg"
//	@Security		Bearer
//
//	@Router			/communication/{id} [put]
func MakeUpdateTemplateCommsEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(UpdateCommunicationTemplateByIDRequest)
		if !ok {
			return nil, ie.Error{
				Code:    http.StatusBadRequest,
				Message: "request has incorrect structure",
			}
		}

		updateTemplateComms := templateservice.UpdateCommunicationTemplateRequest(req)
		templateComms, err := s.UpdateCommunicationTemplateByID(ctx, &updateTemplateComms)
		if err != nil {
			return nil, err
		}

		resp := CommunicationTemplateResponse(*templateComms)
		return resp, nil
	}
}
