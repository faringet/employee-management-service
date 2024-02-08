package endpoints

import (
	"context"
	"net/http"
	"time"

	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

type GetCommunicationTemplateByIDWithEntity struct {
	ID int
}

type CommunicationTemplateWithEntityResponse struct {
	ID                 int            `json:"id"`
	OwnerEntityID      int            `json:"owner_entity_id"`
	HeaderLogoID       int            `json:"header_logo_id"`
	ReminderDaysID     int            `json:"reminder_days_id"`
	IsSendReport       *bool          `json:"is_send_report"`
	Name               *string        `json:"name"`
	Description        *string        `json:"description"`
	TimeSendReport     *time.Time     `json:"time_send_report"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	CreatedBy          *int           `json:"created_by"`
	UpdatedBy          *int           `json:"updated_by"`
	EntityInfoResponse EntityResponse `json:"entity"`
}

type EntityResponse struct {
	ID                           int        `json:"id"`
	WorkspaceID                  int        `json:"workspace_id"`
	CompanyStatusID              int        `json:"company_status_id"`
	CustomerStatusID             int        `json:"customer_status_id"`
	OrganizationSizeCategoriesID int        `json:"organization_size_categories_id"`
	Name                         *string    `json:"name"`
	BoldBiSiteName               *string    `json:"bold_bi_site_name"`
	Details                      *string    `json:"details"`
	ImportLock                   *bool      `json:"import_lock"`
	CreatedAt                    *time.Time `json:"created_at"`
	UpdatedAt                    *time.Time `json:"updated_at"`
	CreatedBy                    *int       `json:"created_by"`
	UpdatedBy                    *int       `json:"updated_by"`
}

// Profile godoc
// Get communication template by id with entity
//
//	@Summary		Get communication template
//	@Description	Get communication template with entity record by id
//	@Tags			communication
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int										true	"ID"
//
//	@Success		200	{object}	CommunicationTemplateWithEntityResponse	"json with empty response"
//	@Failure		400	{object}	SwaggerSimpleError						"json with error msg"
//	@Failure		401	{object}	SwaggerError							"json with error msg"
//	@Failure		403	{object}	SwaggerSimpleError						"json with error msg"
//	@Failure		404	{object}	SwaggerSimpleError						"json with error msg"
//	@Failure		422	{object}	SwaggerError							"json with error msg"
//	@Failure		500	{object}	SwaggerSimpleError						"json with error msg"
//	@Security		Bearer
//
//	@Router			/communication/{id}/entity [get]
func MakeGetCommunicationTemplateByIDWithEntityEndpoint(s templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetCommunicationTemplateByIDWithEntity)
		if !ok {
			return nil, ie.Error{
				Code:    http.StatusBadRequest,
				Message: "request has incorrect structure",
			}
		}

		getCommunicationTemplateReq := templateservice.GetCommunicationTemplateByIDWithEntityRequest(req)
		templateComms, err := s.GetCommunicationTemplateByIDWithEntity(ctx, &getCommunicationTemplateReq)
		if err != nil {
			return nil, err
		}

		return &CommunicationTemplateWithEntityResponse{
			ID:                 templateComms.ID,
			OwnerEntityID:      templateComms.OwnerEntityID,
			HeaderLogoID:       templateComms.HeaderLogoID,
			ReminderDaysID:     templateComms.ReminderDaysID,
			IsSendReport:       templateComms.IsSendReport,
			Name:               templateComms.Name,
			Description:        templateComms.Description,
			TimeSendReport:     templateComms.TimeSendReport,
			CreatedAt:          templateComms.CreatedAt,
			UpdatedAt:          templateComms.UpdatedAt,
			CreatedBy:          templateComms.CreatedBy,
			UpdatedBy:          templateComms.UpdatedBy,
			EntityInfoResponse: EntityResponse(templateComms.EntityInfo),
		}, nil
	}
}
