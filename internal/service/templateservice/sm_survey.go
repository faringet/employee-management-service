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

type SmSurvey interface {
	CreateSmSurvey(ctx context.Context, input BaseSmSurveyDTO) (*BaseIdResponse, error)
	CreateSmSurveyRange(ctx context.Context, input []BaseSmSurveyDTO) error
	DeleteSmSurveyByID(ctx context.Context, input BaseIdRequest) error
	GetSmSurveyByID(ctx context.Context, req BaseIdRequest) (*BaseSmSurveyDTO, error)
	GetSmSurvey(ctx context.Context, req GetSmSurveyRequest) (*GetSmSurveyResponse, error)
	UpdateSmSurveyByID(ctx context.Context, input BaseSmSurveyDTO) (*BaseIdResponse, error)
}

type BaseSmSurveyDTO struct {
	ModifierEmail            string
	ID                       int32      `json:"id"`
	Name                     *string    `json:"name"`
	UpdatedAt                *time.Time `json:"updated_at"`
	CreatedAt                *time.Time `json:"created_at"`
	UpdatedBy                *int32     `json:"updated_by"`
	CreatedBy                *int32     `json:"created_by"`
	SurveyDateStart          *time.Time `json:"survey_date_start"`
	SurveyDateEnd            *time.Time `json:"survey_date_end"`
	StatusID                 int32      `json:"status_id"`
	ProjectID                int32      `json:"project_id"`
	MilestoneDay             *int32     `json:"milestone_day"`
	DateAttributeMilestoneID *int32     `json:"date_attribute_milestone_id"`
	TimestartMilestone       *time.Time `json:"timestart_milestone"`

	SmSurveyStatus *BaseSmSurveyStatusDTO `json:"sm_survey_status"`
	SmProject      *BaseSmProjectDTO      `json:"sm_project"`
}

type GetSmSurveyRequest struct {
	PaginationRequest
}

type CreateSmSurveyRangeDTO struct {
	CreatorEmail string
	SmSurvey     []BaseSmSurveyDTO
}

type GetSmSurveyResponse struct {
	PaginationResponse
	Data []BaseSmSurveyDTO `json:"data"`
}

func (s *service) CreateSmSurvey(ctx context.Context, input BaseSmSurveyDTO) (*BaseIdResponse, error) {
	accId := int32(1)

	req := model.SmSurvey{
		CreatedBy:                &accId,
		UpdatedBy:                &accId,
		Name:                     input.Name,
		SurveyDateStart:          input.SurveyDateStart,
		SurveyDateEnd:            input.SurveyDateEnd,
		StatusID:                 input.StatusID,
		ProjectID:                input.ProjectID,
		MilestoneDay:             input.MilestoneDay,
		DateAttributeMilestoneID: input.DateAttributeMilestoneID,
		TimestartMilestone:       input.TimestartMilestone,
	}

	res, err := s.repo.CreateSmSurvey(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateSmSurveyRange(ctx context.Context, input []BaseSmSurveyDTO) error {

	var rows []model.SmSurvey
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.SmSurvey{
			CreatedBy:                &accId,
			UpdatedBy:                &accId,
			Name:                     row.Name,
			SurveyDateStart:          row.SurveyDateStart,
			SurveyDateEnd:            row.SurveyDateEnd,
			StatusID:                 row.StatusID,
			ProjectID:                row.ProjectID,
			MilestoneDay:             row.MilestoneDay,
			DateAttributeMilestoneID: row.DateAttributeMilestoneID,
			TimestartMilestone:       row.TimestartMilestone,
		})
	}

	err := s.repo.CreateSmSurveyRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetSmSurveyByID(ctx context.Context, req BaseIdRequest) (*BaseSmSurveyDTO, error) {
	row, err := s.repo.GetSmSurveyByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseSmSurveyDTO{
		ID:                       row.ID,
		Name:                     row.Name,
		UpdatedAt:                row.UpdatedAt,
		CreatedAt:                row.CreatedAt,
		UpdatedBy:                row.UpdatedBy,
		CreatedBy:                row.CreatedBy,
		SurveyDateStart:          row.SurveyDateStart,
		SurveyDateEnd:            row.SurveyDateEnd,
		StatusID:                 row.StatusID,
		ProjectID:                row.ProjectID,
		MilestoneDay:             row.MilestoneDay,
		DateAttributeMilestoneID: row.DateAttributeMilestoneID,
		TimestartMilestone:       row.TimestartMilestone,

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
		SmSurveyStatus: &BaseSmSurveyStatusDTO{
			ID:          row.SmSurveyStatus.ID,
			Name:        row.SmSurveyStatus.Name,
			Code:        row.SmSurveyStatus.Code,
			Description: row.SmSurveyStatus.Description,
			UpdatedAt:   row.SmSurveyStatus.UpdatedAt,
			CreatedBy:   row.SmSurveyStatus.CreatedBy,
			CreatedAt:   row.SmSurveyStatus.CreatedAt,
			UpdatedBy:   row.SmSurveyStatus.UpdatedBy,
		},
	}

	return tq, nil
}

func (s *service) GetSmSurvey(ctx context.Context, req GetSmSurveyRequest) (*GetSmSurveyResponse, error) {

	input := repomodel.GetSmSurveyRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetSmSurvey(ctx, input)
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

	res := []BaseSmSurveyDTO{}
	for _, row := range rows {
		tq := BaseSmSurveyDTO{
			ID:                       row.ID,
			Name:                     row.Name,
			UpdatedAt:                row.UpdatedAt,
			CreatedAt:                row.CreatedAt,
			UpdatedBy:                row.UpdatedBy,
			CreatedBy:                row.CreatedBy,
			SurveyDateStart:          row.SurveyDateStart,
			SurveyDateEnd:            row.SurveyDateEnd,
			StatusID:                 row.StatusID,
			ProjectID:                row.ProjectID,
			MilestoneDay:             row.MilestoneDay,
			DateAttributeMilestoneID: row.DateAttributeMilestoneID,
			TimestartMilestone:       row.TimestartMilestone,

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
			SmSurveyStatus: &BaseSmSurveyStatusDTO{
				ID:          row.SmSurveyStatus.ID,
				Name:        row.SmSurveyStatus.Name,
				Code:        row.SmSurveyStatus.Code,
				Description: row.SmSurveyStatus.Description,
				UpdatedAt:   row.SmSurveyStatus.UpdatedAt,
				CreatedBy:   row.SmSurveyStatus.CreatedBy,
				CreatedAt:   row.SmSurveyStatus.CreatedAt,
				UpdatedBy:   row.SmSurveyStatus.UpdatedBy,
			},
		}

		res = append(res, tq)
	}

	count, err := s.repo.CountSmSurvey(ctx)
	if err != nil {
		s.logger.Error("count SmSurvey", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetSmSurveyResponse{
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

func (s *service) UpdateSmSurveyByID(ctx context.Context, input BaseSmSurveyDTO) (*BaseIdResponse, error) {

	var accId int32 = 1

	res, err := s.repo.UpdateSmSurveyByID(ctx, model.SmSurvey{
		ID:                       input.ID,
		UpdatedBy:                &accId,
		Name:                     input.Name,
		SurveyDateStart:          input.SurveyDateStart,
		SurveyDateEnd:            input.SurveyDateEnd,
		StatusID:                 input.StatusID,
		ProjectID:                input.ProjectID,
		MilestoneDay:             input.MilestoneDay,
		DateAttributeMilestoneID: input.DateAttributeMilestoneID,
		TimestartMilestone:       input.TimestartMilestone,
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

func (s *service) DeleteSmSurveyByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteSmSurveyByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
