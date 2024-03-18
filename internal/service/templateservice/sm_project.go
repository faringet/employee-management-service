package templateservice

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
	"go.uber.org/zap"
)

type SmProject interface {
	CreateSmProject(ctx context.Context, input BaseSmProjectDTO) (*BaseIdResponse, error)
	CreateSmProjectRange(ctx context.Context, input []BaseSmProjectDTO) error
	DeleteSmProjectByID(ctx context.Context, input BaseIdRequest) error
	GetSmProjectByID(ctx context.Context, req BaseIdRequest) (*BaseSmProjectDTO, error)
	GetSmProject(ctx context.Context, req GetSmProjectRequest) (*GetSmProjectResponse, error)
	UpdateSmProjectByID(ctx context.Context, input BaseSmProjectDTO) (*BaseIdResponse, error)
}

type BaseSmProjectDTO struct {
	ModifierEmail            string
	ID                       int32      `json:"id"`
	Name                     *string    `json:"name"`
	ProjecttypeID            int32      `json:"projecttype_id"`
	MinResponses             *int32     `json:"min_responses"`
	DateEnd                  *time.Time `json:"date_end"`
	AccessLink               *string    `json:"access_link"`
	UpdatedAt                *time.Time `json:"updated_at"`
	CreatedAt                *time.Time `json:"created_at"`
	UpdatedBy                *int32     `json:"updated_by"`
	CreatedBy                *int32     `json:"created_by"`
	DateAttributeMilestoneID *int32     `json:"date_attribute_milestone_id"`

	SmProjectType       *BaseSmProjectTypeDTO        `json:"sm_project_type"`
	Attributes          *BaseAttributesDTO           `json:"attributes"`
	SmAttributeTriggers []BaseSmAttributeTriggersDTO `json:"sm_attribute_triggers"`
	SmSurvey            []BaseSmSurveyDTO            `json:"sm_survey"`
}

type GetSmProjectRequest struct {
	PaginationRequest
}

type CreateSmProjectRangeDTO struct {
	CreatorEmail string
	SmProject    []BaseSmProjectDTO
}

type GetSmProjectResponse struct {
	PaginationResponse
	Data []BaseSmProjectDTO `json:"data"`
}

func (s *service) CreateSmProject(ctx context.Context, input BaseSmProjectDTO) (*BaseIdResponse, error) {

	accId := int32(1)

	req := model.SmProject{
		CreatedBy:                &accId,
		UpdatedBy:                &accId,
		Name:                     input.Name,
		ProjecttypeID:            input.ProjecttypeID,
		MinResponses:             input.MinResponses,
		DateEnd:                  input.DateEnd,
		AccessLink:               input.AccessLink,
		DateAttributeMilestoneID: input.DateAttributeMilestoneID,
	}

	res, err := s.repo.CreateSmProject(ctx, req)
	if err != nil {
		s.logger.Error("create SmProject", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateSmProjectRange(ctx context.Context, input []BaseSmProjectDTO) error {

	var rows []model.SmProject
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.SmProject{
			CreatedBy:                &accId,
			UpdatedBy:                &accId,
			Name:                     row.Name,
			ProjecttypeID:            row.ProjecttypeID,
			MinResponses:             row.MinResponses,
			DateEnd:                  row.DateEnd,
			AccessLink:               row.AccessLink,
			DateAttributeMilestoneID: row.DateAttributeMilestoneID,
		})
	}

	err := s.repo.CreateSmProjectRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetSmProjectByID(ctx context.Context, req BaseIdRequest) (*BaseSmProjectDTO, error) {
	row, err := s.repo.GetSmProjectByID(ctx, repomodel.BaseIdRequest{
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
			//  s.logger.Error(, zap.Error(err))
			return nil, ie.ErrInternalServerError
		}
	}

	tq := BaseSmProjectDTO{
		ID:                       row.ID,
		UpdatedAt:                row.UpdatedAt,
		CreatedAt:                row.CreatedAt,
		UpdatedBy:                row.UpdatedBy,
		CreatedBy:                row.CreatedBy,
		Name:                     row.Name,
		ProjecttypeID:            row.ProjecttypeID,
		MinResponses:             row.MinResponses,
		DateEnd:                  row.DateEnd,
		AccessLink:               row.AccessLink,
		DateAttributeMilestoneID: row.DateAttributeMilestoneID,
	}

	if row.Attributes != nil {
		tq.Attributes = &BaseAttributesDTO{
			ID:        row.Attributes.ID,
			Name:      row.Attributes.Name,
			IsDate:    row.Attributes.IsDate,
			CreatedAt: row.Attributes.CreatedAt,
			UpdatedAt: row.Attributes.UpdatedAt,
			CreatedBy: row.Attributes.CreatedBy,
			UpdatedBy: row.Attributes.UpdatedBy,
		}
	}

	if row.SmProjectType != nil {
		tq.SmProjectType = &BaseSmProjectTypeDTO{
			ID:          row.SmProjectType.ID,
			Name:        row.SmProjectType.Name,
			Code:        row.SmProjectType.Code,
			Description: row.SmProjectType.Description,
			UpdatedAt:   row.SmProjectType.UpdatedAt,
			CreatedAt:   row.SmProjectType.CreatedAt,
			UpdatedBy:   row.SmProjectType.UpdatedBy,
			CreatedBy:   row.SmProjectType.CreatedBy,
		}
	}
	sm_survey := []BaseSmSurveyDTO{}
	for _, r := range row.SmSurvey {
		sm_survey = append(sm_survey, BaseSmSurveyDTO{
			ID:                       r.ID,
			Name:                     r.Name,
			UpdatedAt:                r.UpdatedAt,
			CreatedAt:                r.CreatedAt,
			UpdatedBy:                r.UpdatedBy,
			CreatedBy:                r.CreatedBy,
			SurveyDateStart:          r.SurveyDateStart,
			SurveyDateEnd:            r.SurveyDateEnd,
			StatusID:                 r.StatusID,
			ProjectID:                r.ProjectID,
			MilestoneDay:             r.MilestoneDay,
			DateAttributeMilestoneID: r.DateAttributeMilestoneID,
			TimestartMilestone:       r.TimestartMilestone,
		})
	}
	tq.SmSurvey = sm_survey
	sm_attribute_triggers := []BaseSmAttributeTriggersDTO{}
	for _, r := range row.SmAttributeTriggers {
		sm_attribute_triggers = append(sm_attribute_triggers, BaseSmAttributeTriggersDTO{
			ID:          r.ID,
			ProjectID:   r.ProjectID,
			AttributeID: r.AttributeID,
			Value:       r.Value,
			UpdatedAt:   r.UpdatedAt,
			CreatedAt:   r.CreatedAt,
			UpdatedBy:   r.UpdatedBy,
			CreatedBy:   r.CreatedBy,
		})
	}
	tq.SmAttributeTriggers = sm_attribute_triggers

	return &tq, nil
}

func (s *service) GetSmProject(ctx context.Context, req GetSmProjectRequest) (*GetSmProjectResponse, error) {

	input := repomodel.GetSmProjectRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	log.Println(input)

	rows, err := s.repo.GetSmProject(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, ie.ErrDBRecordNotFound):
			return nil, ie.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:

			//  s.logger.Error(, zap.Error(err))
			return nil, ie.ErrInternalServerError
		}
	}
	res := []BaseSmProjectDTO{}

	for _, row := range rows {
		tq := BaseSmProjectDTO{
			ID:                       row.ID,
			UpdatedAt:                row.UpdatedAt,
			CreatedAt:                row.CreatedAt,
			UpdatedBy:                row.UpdatedBy,
			CreatedBy:                row.CreatedBy,
			Name:                     row.Name,
			ProjecttypeID:            row.ProjecttypeID,
			MinResponses:             row.MinResponses,
			DateEnd:                  row.DateEnd,
			AccessLink:               row.AccessLink,
			DateAttributeMilestoneID: row.DateAttributeMilestoneID,
		}

		// список справочников
		// начало справочника #1
		if row.Attributes != nil {
			tq.Attributes = &BaseAttributesDTO{
				ID:        row.Attributes.ID,
				Name:      row.Attributes.Name,
				IsDate:    row.Attributes.IsDate,
				CreatedAt: row.Attributes.CreatedAt,
				UpdatedAt: row.Attributes.UpdatedAt,
				CreatedBy: row.Attributes.CreatedBy,
				UpdatedBy: row.Attributes.UpdatedBy,
			}
		}

		sm_survey := []BaseSmSurveyDTO{}
		for _, r := range row.SmSurvey {
			sm_survey = append(sm_survey, BaseSmSurveyDTO{
				ID:                       r.ID,
				Name:                     r.Name,
				UpdatedAt:                r.UpdatedAt,
				CreatedAt:                r.CreatedAt,
				UpdatedBy:                r.UpdatedBy,
				CreatedBy:                r.CreatedBy,
				SurveyDateStart:          r.SurveyDateStart,
				SurveyDateEnd:            r.SurveyDateEnd,
				StatusID:                 r.StatusID,
				ProjectID:                r.ProjectID,
				MilestoneDay:             r.MilestoneDay,
				DateAttributeMilestoneID: r.DateAttributeMilestoneID,
				TimestartMilestone:       r.TimestartMilestone,
			})
		}
		tq.SmSurvey = sm_survey
		sm_attribute_triggers := []BaseSmAttributeTriggersDTO{}
		for _, r := range row.SmAttributeTriggers {
			sm_attribute_triggers = append(sm_attribute_triggers, BaseSmAttributeTriggersDTO{
				ID:          r.ID,
				ProjectID:   r.ProjectID,
				AttributeID: r.AttributeID,
				Value:       r.Value,
				UpdatedAt:   r.UpdatedAt,
				CreatedAt:   r.CreatedAt,
				UpdatedBy:   r.UpdatedBy,
				CreatedBy:   r.CreatedBy,
			})
		}
		tq.SmAttributeTriggers = sm_attribute_triggers

		// конец mtm #1

		res = append(res, tq)

	}

	count, err := s.repo.CountSmProject(ctx)
	if err != nil {
		s.logger.Error("count SmProject", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetSmProjectResponse{
		Data: res,
		PaginationResponse: PaginationResponse{
			ActivePage:   activePage,
			TotalCount:   *count,
			CountPerPage: req.Limit,
			TotalPages:   totalPages,
		},
	}
	log.Println(page_res)

	return page_res, nil
}

func (s *service) UpdateSmProjectByID(ctx context.Context, input BaseSmProjectDTO) (*BaseIdResponse, error) {

	var accId int32 = 1

	res, err := s.repo.UpdateSmProjectByID(ctx, model.SmProject{
		ID:        int32(input.ID),
		UpdatedBy: &accId,

		Name:                     input.Name,
		ProjecttypeID:            input.ProjecttypeID,
		MinResponses:             input.MinResponses,
		DateEnd:                  input.DateEnd,
		AccessLink:               input.AccessLink,
		DateAttributeMilestoneID: input.DateAttributeMilestoneID,
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

func (s *service) DeleteSmProjectByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteSmProjectByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
