package templateservice

import (
	"context"
	"errors"
	"github.com/engagerocketco/go-common/ns"
	"github.com/engagerocketco/templates-api-svc/internal/repository"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/service/entityservice"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type CommunicationTemplateService interface {
	CreateCommunicationTemplate(ctx context.Context, req *CreateCommunicationTemplateRequest) (*CommunicationTemplateResponse, error)
	GetCommunicationTemplateByID(ctx context.Context, req *GetCommunicationTemplateByIDRequest) (*CommunicationTemplateResponse, error)
	GetCommunicationTemplatesByEntityID(ctx context.Context, req *GetCommunicationTemplatesByEntityIDRequest) ([]CommunicationTemplateResponse, error)
	GetCommunicationTemplateByIDWithEntity(ctx context.Context, req *GetCommunicationTemplateByIDWithEntityRequest) (*CommunicationTemplateWithEntityResponse, error)
	UpdateCommunicationTemplateByID(ctx context.Context, req *UpdateCommunicationTemplateRequest) (*CommunicationTemplateResponse, error)
	DeleteCommunicationTemplateByID(ctx context.Context, req *DeleteCommunicationTemplateByIDRequest) error
}

type GetCommunicationTemplateByIDRequest struct {
	ID int
}

type DeleteCommunicationTemplateByIDRequest struct {
	ID int
}

type CreateCommunicationTemplateRequest struct {
	OwnerEntityID  int
	HeaderLogoID   int
	ReminderDaysID int
	IsSendReport   *bool
	Name           *string
	Description    *string
	TimeSendReport *time.Time
	CreatedBy      *int
	UpdatedBy      *int
}

type UpdateCommunicationTemplateRequest struct {
	ID             int
	OwnerEntityID  int
	HeaderLogoID   int
	ReminderDaysID int
	IsSendReport   *bool
	Name           *string
	Description    *string
	TimeSendReport *time.Time
	UpdatedBy      *int
}

type GetCommunicationTemplatesByEntityIDRequest struct {
	ID int
}

type GetCommunicationTemplateByIDWithEntityRequest struct {
	ID int
}

type CommunicationTemplateResponse struct {
	ID             int
	OwnerEntityID  int
	HeaderLogoID   int
	ReminderDaysID int
	IsSendReport   *bool
	Name           *string
	Description    *string
	TimeSendReport *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	CreatedBy      *int
	UpdatedBy      *int
}

type CommunicationTemplateWithEntityResponse struct {
	ID             int
	OwnerEntityID  int
	HeaderLogoID   int
	ReminderDaysID int
	IsSendReport   *bool
	Name           *string
	Description    *string
	TimeSendReport *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	CreatedBy      *int
	UpdatedBy      *int
	EntityInfo     entityservice.Entity
}

func (s *service) CreateCommunicationTemplate(ctx context.Context, createCommunicationTemplate *CreateCommunicationTemplateRequest) (*CommunicationTemplateResponse, error) {
	req := repomodel.CreateCommunicationTemplate(*createCommunicationTemplate)
	templComms, err := s.repo.CreateCommunicationTemplate(ctx, &req)
	if err != nil {
		s.logger.Error("failed to create communication template record", zap.Error(err), zap.String("name", *req.Name))
		return nil, ie.ErrInternalServerError
	}

	res := CommunicationTemplateResponse(*templComms)
	return &res, nil
}

func (s *service) GetCommunicationTemplateByID(ctx context.Context, req *GetCommunicationTemplateByIDRequest) (*CommunicationTemplateResponse, error) {
	templComms, err := s.repo.GetCommunicationTemplateByID(ctx, req.ID)
	if err != nil {
		s.logger.Error("failed to get communication template record by id", zap.Error(err), zap.Int("id", req.ID))
		if errors.Is(err, repository.ErrCommunicationTemplateNotFound) {
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: repository.ErrCommunicationTemplateNotFound.Error(),
			}
		}

		return nil, ie.ErrInternalServerError
	}

	res := CommunicationTemplateResponse(*templComms)
	return &res, nil
}

func (s *service) GetCommunicationTemplatesByEntityID(ctx context.Context, req *GetCommunicationTemplatesByEntityIDRequest) ([]CommunicationTemplateResponse, error) {
	templComms, err := s.repo.GetCommunicationTemplatesByEntityID(ctx, req.ID)
	if err != nil {
		s.logger.Error("failed to get communication template records by entity id", zap.Error(err), zap.Int("id", req.ID))
		if errors.Is(err, repository.ErrCommunicationTemplateNotFound) {
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: repository.ErrCommunicationTemplateNotFound.Error(),
			}
		}

		return nil, ie.ErrInternalServerError
	}

	var res []CommunicationTemplateResponse
	for _, template := range templComms {
		res = append(res, CommunicationTemplateResponse(template))
	}

	return res, nil
}

func (s *service) GetCommunicationTemplateByIDWithEntity(ctx context.Context, req *GetCommunicationTemplateByIDWithEntityRequest) (*CommunicationTemplateWithEntityResponse, error) {
	templComms, err := s.repo.GetCommunicationTemplateByID(ctx, req.ID)
	if err != nil {
		s.logger.Error("failed to get communication template with entity record by id", zap.Error(err), zap.Int("id", req.ID))
		if errors.Is(err, repository.ErrCommunicationTemplateNotFound) {
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: repository.ErrCommunicationTemplateNotFound.Error(),
			}
		}

		return nil, ie.ErrInternalServerError
	}

	entityInfo, err := s.natsService.GetEntityByID(ctx, templComms.OwnerEntityID)
	if err != nil {
		s.logger.Error("failed to get entity record from nats", zap.Error(err), zap.Int("id", req.ID))
		if errors.Is(err, ns.ErrEntityNotFound) {
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: repository.ErrCommunicationTemplateNotFound.Error(),
			}
		}

		return nil, ie.ErrInternalServerError
	}

	return &CommunicationTemplateWithEntityResponse{
		ID:             templComms.ID,
		OwnerEntityID:  templComms.OwnerEntityID,
		HeaderLogoID:   templComms.HeaderLogoID,
		ReminderDaysID: templComms.ReminderDaysID,
		IsSendReport:   templComms.IsSendReport,
		Name:           templComms.Name,
		Description:    templComms.Description,
		TimeSendReport: templComms.TimeSendReport,
		CreatedAt:      templComms.CreatedAt,
		UpdatedAt:      templComms.UpdatedAt,
		CreatedBy:      templComms.CreatedBy,
		UpdatedBy:      templComms.UpdatedBy,
		EntityInfo: entityservice.Entity{
			ID:                           entityInfo.ID,
			WorkspaceID:                  entityInfo.WorkspaceID,
			CompanyStatusID:              *entityInfo.CompanyStatusID,
			CustomerStatusID:             *entityInfo.CustomerStatusID,
			OrganizationSizeCategoriesID: *entityInfo.OrganizationSizeCategoriesID,
			Name:                         entityInfo.Name,
			BoldBiSiteName:               entityInfo.BoldBISiteName,
			Details:                      entityInfo.Details,
			ImportLock:                   entityInfo.ImportLock,
			CreatedAt:                    entityInfo.CreatedAt,
			UpdatedAt:                    entityInfo.UpdatedAt,
			CreatedBy:                    entityInfo.CreatedBy,
			UpdatedBy:                    entityInfo.UpdatedBy,
		},
	}, nil
}

func (s *service) UpdateCommunicationTemplateByID(ctx context.Context, updateCommunicationTemplate *UpdateCommunicationTemplateRequest) (*CommunicationTemplateResponse, error) {
	templComms, err := s.repo.UpdateCommunicationTemplateByID(ctx, &repomodel.UpdateCommunicationTemplate{
		ID:             updateCommunicationTemplate.ID,
		OwnerEntityID:  updateCommunicationTemplate.OwnerEntityID,
		HeaderLogoID:   updateCommunicationTemplate.HeaderLogoID,
		ReminderDaysID: updateCommunicationTemplate.ReminderDaysID,
		IsSendReport:   updateCommunicationTemplate.IsSendReport,
		Name:           updateCommunicationTemplate.Name,
		Description:    updateCommunicationTemplate.Description,
		TimeSendReport: updateCommunicationTemplate.TimeSendReport,
		UpdatedBy:      updateCommunicationTemplate.UpdatedBy,
	})
	if err != nil {
		s.logger.Error("failed to update communication template record by id", zap.Error(err), zap.Int("id", updateCommunicationTemplate.ID))
		if errors.Is(err, repository.ErrCommunicationTemplateNotFound) {
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: repository.ErrCommunicationTemplateNotFound.Error(),
			}
		}

		return nil, ie.ErrInternalServerError
	}

	res := CommunicationTemplateResponse(*templComms)
	return &res, nil
}

func (s *service) DeleteCommunicationTemplateByID(ctx context.Context, req *DeleteCommunicationTemplateByIDRequest) error {
	err := s.repo.DeleteCommunicationTemplateByID(ctx, req.ID)
	if err != nil {
		s.logger.Error("failed to delete communication template record by id", zap.Error(err), zap.Int("id", req.ID))
		if errors.Is(err, repository.ErrCommunicationTemplateNotFound) {
			return ie.Error{
				Code:    http.StatusNotFound,
				Message: repository.ErrCommunicationTemplateNotFound.Error(),
			}
		}

		return ie.ErrInternalServerError
	}

	return nil
}
