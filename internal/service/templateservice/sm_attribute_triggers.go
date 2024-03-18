package templateservice

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"go.uber.org/zap"
)

type SmAttributeTriggers interface {
	CreateSmAttributeTriggers(ctx context.Context, input BaseSmAttributeTriggersDTO) (*BaseIdResponse, error)
	CreateSmAttributeTriggersRange(ctx context.Context, input []BaseSmAttributeTriggersDTO) error
	DeleteSmAttributeTriggersByID(ctx context.Context, input BaseIdRequest) error
	GetSmAttributeTriggersByID(ctx context.Context, req BaseIdRequest) (*BaseSmAttributeTriggersDTO, error)
	GetSmAttributeTriggers(ctx context.Context, req GetSmAttributeTriggersRequest) (*GetSmAttributeTriggersResponse, error)
	UpdateSmAttributeTriggersByID(ctx context.Context, input BaseSmAttributeTriggersDTO) (*BaseIdResponse, error)
}

type BaseSmAttributeTriggersDTO struct {
	ModifierEmail string
	ID            int32              `json:"id"`
	ProjectID     int32              `json:"project_id"`
	AttributeID   int32              `json:"attribute_id"`
	Value         *string            `json:"value"`
	UpdatedAt     *time.Time         `json:"updated_at"`
	CreatedAt     *time.Time         `json:"created_at"`
	UpdatedBy     *int32             `json:"updated_by"`
	CreatedBy     *int32             `json:"created_by"`
	Attributes    *BaseAttributesDTO `json:"attributes"`
	SmProject     *BaseSmProjectDTO  `json:"sm_project"`
}

type GetSmAttributeTriggersRequest struct {
	PaginationRequest
}

type CreateSmAttributeTriggersRangeDTO struct {
	CreatorEmail        string
	SmAttributeTriggers []BaseSmAttributeTriggersDTO
}

type GetSmAttributeTriggersResponse struct {
	PaginationResponse
	Data []BaseSmAttributeTriggersDTO `json:"data"`
}

func (s *service) CreateSmAttributeTriggers(ctx context.Context, input BaseSmAttributeTriggersDTO) (*BaseIdResponse, error) {

	accId := int32(1)

	req := model.SmAttributeTriggers{
		CreatedBy:   &accId,
		UpdatedBy:   &accId,
		ProjectID:   input.ProjectID,
		AttributeID: input.AttributeID,
		Value:       input.Value,
	}

	res, err := s.repo.CreateSmAttributeTriggers(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateSmAttributeTriggersRange(ctx context.Context, input []BaseSmAttributeTriggersDTO) error {

	var rows []model.SmAttributeTriggers
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.SmAttributeTriggers{
			CreatedBy:   &accId,
			UpdatedBy:   &accId,
			ProjectID:   row.ProjectID,
			AttributeID: row.AttributeID,
			Value:       row.Value,
		})
	}

	err := s.repo.CreateSmAttributeTriggersRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetSmAttributeTriggersByID(ctx context.Context, req BaseIdRequest) (*BaseSmAttributeTriggersDTO, error) {
	row, err := s.repo.GetSmAttributeTriggersByID(ctx, repomodel.BaseIdRequest{
		ID: int32(req.ID),
	})
	if err != nil {
		switch {
		case errors.Is(err, ie.ErrDBRecordNotFound):
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			return nil, ie.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}

	}

	tq := &BaseSmAttributeTriggersDTO{
		ID:          row.ID,
		ProjectID:   row.ProjectID,
		AttributeID: row.AttributeID,
		Value:       row.Value,
		UpdatedAt:   row.UpdatedAt,
		CreatedAt:   row.CreatedAt,
		UpdatedBy:   row.UpdatedBy,
		CreatedBy:   row.CreatedBy,
		Attributes: &BaseAttributesDTO{
			ID:        row.Attributes.ID,
			Name:      row.Attributes.Name,
			IsDate:    row.Attributes.IsDate,
			CreatedAt: row.Attributes.CreatedAt,
			UpdatedAt: row.Attributes.UpdatedAt,
			CreatedBy: row.Attributes.CreatedBy,
			UpdatedBy: row.Attributes.UpdatedBy,
		},
		SmProject: &BaseSmProjectDTO{
			ID:                       row.SmProject.ID,
			Name:                     row.SmProject.Name,
			ProjecttypeID:            row.SmProject.ProjecttypeID,
			MinResponses:             row.SmProject.MinResponses,
			DateEnd:                  row.SmProject.DateEnd,
			AccessLink:               row.SmProject.AccessLink,
			UpdatedAt:                row.SmProject.UpdatedAt,
			CreatedAt:                row.SmProject.CreatedAt,
			UpdatedBy:                row.SmProject.UpdatedBy,
			CreatedBy:                row.SmProject.CreatedBy,
			DateAttributeMilestoneID: row.SmProject.DateAttributeMilestoneID,
		},
	}

	return tq, nil
}

func (s *service) GetSmAttributeTriggers(ctx context.Context, req GetSmAttributeTriggersRequest) (*GetSmAttributeTriggersResponse, error) {

	input := repomodel.GetSmAttributeTriggersRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetSmAttributeTriggers(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, ie.ErrDBRecordNotFound):
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			s.logger.Error("get all template questionaries", zap.Error(err))
			return nil, ie.ErrInternalServerError
		}
	}

	res := []BaseSmAttributeTriggersDTO{}
	for _, row := range rows {
		tq := BaseSmAttributeTriggersDTO{
			ID:          row.ID,
			ProjectID:   row.ProjectID,
			AttributeID: row.AttributeID,
			Value:       row.Value,
			UpdatedAt:   row.UpdatedAt,
			CreatedAt:   row.CreatedAt,
			UpdatedBy:   row.UpdatedBy,
			CreatedBy:   row.CreatedBy,
			Attributes: &BaseAttributesDTO{
				ID:        row.Attributes.ID,
				Name:      row.Attributes.Name,
				IsDate:    row.Attributes.IsDate,
				CreatedAt: row.Attributes.CreatedAt,
				UpdatedAt: row.Attributes.UpdatedAt,
				CreatedBy: row.Attributes.CreatedBy,
				UpdatedBy: row.Attributes.UpdatedBy,
			},
			SmProject: &BaseSmProjectDTO{
				ID:                       row.SmProject.ID,
				Name:                     row.SmProject.Name,
				ProjecttypeID:            row.SmProject.ProjecttypeID,
				MinResponses:             row.SmProject.MinResponses,
				DateEnd:                  row.SmProject.DateEnd,
				AccessLink:               row.SmProject.AccessLink,
				UpdatedAt:                row.SmProject.UpdatedAt,
				CreatedAt:                row.SmProject.CreatedAt,
				UpdatedBy:                row.SmProject.UpdatedBy,
				CreatedBy:                row.SmProject.CreatedBy,
				DateAttributeMilestoneID: row.SmProject.DateAttributeMilestoneID,
			},
		}

		res = append(res, tq)
	}

	count, err := s.repo.CountSmAttributeTriggers(ctx)
	if err != nil {
		s.logger.Error("count SmAttributeTriggers", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetSmAttributeTriggersResponse{
		Data: res,
		PaginationResponse: PaginationResponse{
			ActivePage:   activePage,
			TotalCount:   *count,
			CountPerPage: req.Limit,
			TotalPages:   totalPages,
		},
	}
	return page_res, nil
}

func (s *service) UpdateSmAttributeTriggersByID(ctx context.Context, input BaseSmAttributeTriggersDTO) (*BaseIdResponse, error) {

	var accId int32 = 1

	res, err := s.repo.UpdateSmAttributeTriggersByID(ctx, model.SmAttributeTriggers{
		ID:          input.ID,
		UpdatedBy:   &accId,
		ProjectID:   input.ProjectID,
		AttributeID: input.AttributeID,
		Value:       input.Value,
	})
	if err != nil {
		switch {
		case errors.Is(err, ie.ErrDBRecordNotFound):
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			return nil, ie.ErrInternalServerError
		}
	}

	e := BaseIdResponse(*res)

	return &e, nil
}

func (s *service) DeleteSmAttributeTriggersByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteSmAttributeTriggersByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
