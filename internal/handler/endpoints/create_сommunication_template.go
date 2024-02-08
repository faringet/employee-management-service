package endpoints

import (
	"context"
	"net/http"
	"time"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type CreateCommunicationTemplateRequest struct {
	OwnerEntityID  int        `json:"owner_entity_id" validate:"required"`
	HeaderLogoID   int        `json:"header_logo_id" validate:"required"`
	ReminderDaysID int        `json:"reminder_days_id" validate:"required"`
	IsSendReport   *bool      `json:"is_send_report" validate:"required"`
	Name           *string    `json:"name" validate:"required"`
	Description    *string    `json:"description" validate:"required"`
	TimeSendReport *time.Time `json:"time_send_report" validate:"required"`
	CreatedBy      *int       `json:"created_by" validate:"required"`
	UpdatedBy      *int       `json:"updated_by" validate:"required"`
}

type CommunicationTemplateResponse struct {
	ID             int        `json:"id"`
	OwnerEntityID  int        `json:"owner_entity_id"`
	HeaderLogoID   int        `json:"header_logo_id"`
	ReminderDaysID int        `json:"reminder_days_id"`
	IsSendReport   *bool      `json:"is_send_report"`
	Name           *string    `json:"name"`
	Description    *string    `json:"description"`
	TimeSendReport *time.Time `json:"time_send_report"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	CreatedBy      *int       `json:"created_by"`
	UpdatedBy      *int       `json:"updated_by"`
}

// Profile godoc
// Create communication template
//
//	@Summary		Create communication template
//	@Description	Create communication template record
//	@Tags			communication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateCommunicationTemplateRequest	true	"Create request"
//
//	@Success		200		{object}	CommunicationTemplateResponse
//	@Failure		400		{object}	SwaggerSimpleError	"json with error msg"
//	@Failure		401		{object}	SwaggerError		"json with error msg"
//	@Failure		422		{object}	SwaggerError		"json with error msg"
//	@Failure		403		{object}	SwaggerSimpleError	"json with error msg"
//	@Failure		500		{object}	SwaggerSimpleError	"json with error msg"
//	@Security		Bearer
//
//	@Router			/communication [post]
func MakeCreateCommunicationTemplateEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateCommunicationTemplateRequest)
		if !ok {
			return nil, ie.Error{
				Code:    http.StatusBadRequest,
				Message: "request has incorrect structure",
			}
		}

		createCommunicationTemplateReq := templateservice.CreateCommunicationTemplateRequest(req)
		templateComms, err := s.CreateCommunicationTemplate(ctx, &createCommunicationTemplateReq)
		if err != nil {
			return nil, err
		}

		resp := CommunicationTemplateResponse(*templateComms)
		return resp, nil
	}
}
