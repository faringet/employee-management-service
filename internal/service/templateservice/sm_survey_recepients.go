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

type SmSurveyRecepients interface {
	CreateSmSurveyRecepients(ctx context.Context, input BaseSmSurveyRecepientsDTO) (*BaseIdResponse, error)
	CreateSmSurveyRecepientsRange(ctx context.Context, input []BaseSmSurveyRecepientsDTO) error
	DeleteSmSurveyRecepientsByID(ctx context.Context, input BaseIdRequest) error
	GetSmSurveyRecepientsByID(ctx context.Context, req BaseIdRequest) (*BaseSmSurveyRecepientsDTO, error)
	GetSmSurveyRecepients(ctx context.Context, req GetSmSurveyRecepientsRequest) (*GetSmSurveyRecepientsResponse, error)
	UpdateSmSurveyRecepientsByID(ctx context.Context, input BaseSmSurveyRecepientsDTO) (*BaseIdResponse, error)
}

type BaseSmSurveyRecepientsDTO struct {
	ModifierEmail     string
	ID                int32             `json:"id"`
	SurveyID          int32             `json:"survey_id"`
	EmployeeID        *int32            `json:"employee_id"`
	Automatical       *bool             `json:"automatical"`
	AccessCode        *string           `json:"access_code"`
	UpdatedAt         *time.Time        `json:"updated_at"`
	CreatedBy         *int32            `json:"created_by"`
	CreatedAt         *time.Time        `json:"created_at"`
	UpdatedBy         *int32            `json:"updated_by"`
	MilestoneDateSend *time.Time        `json:"milestone_date_send"`
	Answered          *bool             `json:"answered"`
	Answers           *string           `json:"answers"`
	Employees         *BaseEmployeesDTO `json:"employees"`
	SmSurvey          *BaseSmSurveyDTO  `json:"sm_survey"`
}

type GetSmSurveyRecepientsRequest struct {
	PaginationRequest
}

type CreateSmSurveyRecepientsRangeDTO struct {
	CreatorEmail       string
	SmSurveyRecepients []BaseSmSurveyRecepientsDTO
}

type GetSmSurveyRecepientsResponse struct {
	PaginationResponse
	Data []BaseSmSurveyRecepientsDTO `json:"data"`
}

func (s *service) CreateSmSurveyRecepients(ctx context.Context, input BaseSmSurveyRecepientsDTO) (*BaseIdResponse, error) {

	accId := int32(1)

	req := model.SmSurveyRecepients{
		CreatedBy:         &accId,
		UpdatedBy:         &accId,
		SurveyID:          input.SurveyID,
		EmployeeID:        input.EmployeeID,
		Automatical:       input.Automatical,
		AccessCode:        input.AccessCode,
		MilestoneDateSend: input.MilestoneDateSend,
		Answered:          input.Answered,
		Answers:           input.Answers,
	}

	res, err := s.repo.CreateSmSurveyRecepients(ctx, req)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", req))
		return nil, ie.ErrInternalServerError
	}

	e := &BaseIdResponse{
		ID: res.ID,
	}

	return e, nil
}

func (s *service) CreateSmSurveyRecepientsRange(ctx context.Context, input []BaseSmSurveyRecepientsDTO) error {

	var rows []model.SmSurveyRecepients
	var accId int32 = 1

	for _, row := range input {
		rows = append(rows, model.SmSurveyRecepients{
			CreatedBy:         &accId,
			UpdatedBy:         &accId,
			SurveyID:          row.SurveyID,
			EmployeeID:        row.EmployeeID,
			Automatical:       row.Automatical,
			AccessCode:        row.AccessCode,
			MilestoneDateSend: row.MilestoneDateSend,
			Answered:          row.Answered,
			Answers:           row.Answers,
		})
	}

	err := s.repo.CreateSmSurveyRecepientsRange(ctx, rows)
	if err != nil {
		s.logger.Error("create template questionary", zap.Error(err), zap.Any("create request", rows))
		return ie.ErrInternalServerError
	}

	return nil
}

func (s *service) GetSmSurveyRecepientsByID(ctx context.Context, req BaseIdRequest) (*BaseSmSurveyRecepientsDTO, error) {
	row, err := s.repo.GetSmSurveyRecepientsByID(ctx, repomodel.BaseIdRequest{
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

	tq := &BaseSmSurveyRecepientsDTO{
		ID:                row.ID,
		SurveyID:          row.SurveyID,
		EmployeeID:        row.EmployeeID,
		Automatical:       row.Automatical,
		AccessCode:        row.AccessCode,
		UpdatedAt:         row.UpdatedAt,
		CreatedBy:         row.CreatedBy,
		CreatedAt:         row.CreatedAt,
		UpdatedBy:         row.UpdatedBy,
		MilestoneDateSend: row.MilestoneDateSend,
		Answered:          row.Answered,
		Answers:           row.Answers,
		SmSurvey: &BaseSmSurveyDTO{
			ID:                       row.SmSurvey.ID,
			Name:                     row.SmSurvey.Name,
			UpdatedAt:                row.SmSurvey.UpdatedAt,
			CreatedAt:                row.SmSurvey.CreatedAt,
			UpdatedBy:                row.SmSurvey.UpdatedBy,
			CreatedBy:                row.SmSurvey.CreatedBy,
			SurveyDateStart:          row.SmSurvey.SurveyDateStart,
			SurveyDateEnd:            row.SmSurvey.SurveyDateEnd,
			StatusID:                 row.SmSurvey.StatusID,
			ProjectID:                row.SmSurvey.ProjectID,
			MilestoneDay:             row.SmSurvey.MilestoneDay,
			DateAttributeMilestoneID: row.SmSurvey.DateAttributeMilestoneID,
			TimestartMilestone:       row.SmSurvey.TimestartMilestone,
		},
		Employees: &BaseEmployeesDTO{
			ID:            row.Employees.ID,
			Email:         row.Employees.Email,
			PreferredName: row.Employees.PreferredName,
			FullName:      row.Employees.FullName,
			UniqueID:      row.Employees.UniqueID,
			UpdatedAt:     row.Employees.UpdatedAt,
			CreatedAt:     row.Employees.CreatedAt,
			UpdatedBy:     row.Employees.UpdatedBy,
			CreatedBy:     row.Employees.CreatedBy,
			ManagerEmail:  row.Employees.ManagerEmail,
		},
	}

	return tq, nil
}

func (s *service) GetSmSurveyRecepients(ctx context.Context, req GetSmSurveyRecepientsRequest) (*GetSmSurveyRecepientsResponse, error) {

	input := repomodel.GetSmSurveyRecepientsRequest{
		PaginationRequest: repomodel.PaginationRequest(req.PaginationRequest),
	}

	rows, err := s.repo.GetSmSurveyRecepients(ctx, input)
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

	res := []BaseSmSurveyRecepientsDTO{}
	for _, row := range rows {
		tq := BaseSmSurveyRecepientsDTO{
			ID:                row.ID,
			SurveyID:          row.SurveyID,
			EmployeeID:        row.EmployeeID,
			Automatical:       row.Automatical,
			AccessCode:        row.AccessCode,
			UpdatedAt:         row.UpdatedAt,
			CreatedBy:         row.CreatedBy,
			CreatedAt:         row.CreatedAt,
			UpdatedBy:         row.UpdatedBy,
			MilestoneDateSend: row.MilestoneDateSend,
			Answered:          row.Answered,
			Answers:           row.Answers,
			SmSurvey: &BaseSmSurveyDTO{
				ID:                       row.SmSurvey.ID,
				Name:                     row.SmSurvey.Name,
				UpdatedAt:                row.SmSurvey.UpdatedAt,
				CreatedAt:                row.SmSurvey.CreatedAt,
				UpdatedBy:                row.SmSurvey.UpdatedBy,
				CreatedBy:                row.SmSurvey.CreatedBy,
				SurveyDateStart:          row.SmSurvey.SurveyDateStart,
				SurveyDateEnd:            row.SmSurvey.SurveyDateEnd,
				StatusID:                 row.SmSurvey.StatusID,
				ProjectID:                row.SmSurvey.ProjectID,
				MilestoneDay:             row.SmSurvey.MilestoneDay,
				DateAttributeMilestoneID: row.SmSurvey.DateAttributeMilestoneID,
				TimestartMilestone:       row.SmSurvey.TimestartMilestone,
			},
			Employees: &BaseEmployeesDTO{
				ID:            row.Employees.ID,
				Email:         row.Employees.Email,
				PreferredName: row.Employees.PreferredName,
				FullName:      row.Employees.FullName,
				UniqueID:      row.Employees.UniqueID,
				UpdatedAt:     row.Employees.UpdatedAt,
				CreatedAt:     row.Employees.CreatedAt,
				UpdatedBy:     row.Employees.UpdatedBy,
				CreatedBy:     row.Employees.CreatedBy,
				ManagerEmail:  row.Employees.ManagerEmail,
			},
		}

		res = append(res, tq)
	}

	count, err := s.repo.CountSmSurveyRecepients(ctx)
	if err != nil {
		s.logger.Error("count SmSurveyRecepients", zap.Error(err))
		return nil, ie.ErrInternalServerError
	}

	totalPages := *count / req.Limit
	if *count%req.Limit > 0 {
		totalPages += 1
	}
	activePage := (input.Offset / req.Limit) + 1
	page_res := &GetSmSurveyRecepientsResponse{
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

func (s *service) UpdateSmSurveyRecepientsByID(ctx context.Context, input BaseSmSurveyRecepientsDTO) (*BaseIdResponse, error) {

	var accId int32 = 1

	res, err := s.repo.UpdateSmSurveyRecepientsByID(ctx, model.SmSurveyRecepients{
		ID:                input.ID,
		UpdatedBy:         &accId,
		SurveyID:          input.SurveyID,
		EmployeeID:        input.EmployeeID,
		Automatical:       input.Automatical,
		AccessCode:        input.AccessCode,
		MilestoneDateSend: input.MilestoneDateSend,
		Answered:          input.Answered,
		Answers:           input.Answers,
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

func (s *service) DeleteSmSurveyRecepientsByID(ctx context.Context, input BaseIdRequest) error {
	err := s.repo.DeleteSmSurveyRecepientsByID(ctx, repomodel.BaseIdRequest{
		ID: int32(input.ID),
	})
	if err != nil {
		s.logger.Error("delete template questionary", zap.Error(err), zap.Any("delete id", input.ID))
		return ie.ErrInternalServerError
	}

	return nil
}
