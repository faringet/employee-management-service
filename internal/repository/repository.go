package repository

import (
	"context"
	"errors"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
)

type Repository interface {
	TemplTemplateQuestionary
	CommunicationTemplateRepository
	SurveyTagsRepository
	TemplRecomendedFrequancyRepository
	TemplQuestionaryTagsRepository
}

type SurveyTagsRepository interface {
	CreateSurveyTags(ctx context.Context, input model.SurveyTags) (*repomodel.BaseIdResponse, error)
	CreateSurveyTagsRange(ctx context.Context, inputs []model.SurveyTags) error
	DeleteSurveyTagsByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetSurveyTags(ctx context.Context) ([]model.SurveyTags, error)
	GetSurveyTagsPagination(ctx context.Context, req repomodel.GetSurveyTagsRequest) ([]model.SurveyTags, error)
	GetSurveyTagsByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.SurveyTags, error)
	UpdateSurveyTagsByID(ctx context.Context, input model.SurveyTags) (*repomodel.BaseIdResponse, error)

	CountSurveyTags(ctx context.Context) (*int, error)
}

type TemplRecomendedFrequancyRepository interface {
	CreateTemplRecomendedFrequancy(ctx context.Context, input model.TemplRecomendedFrequancy) (*repomodel.BaseIdResponse, error)
	CreateTemplRecomendedFrequancyRange(ctx context.Context, inputs []model.TemplRecomendedFrequancy) error
	DeleteTemplRecomendedFrequancyByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetTemplRecomendedFrequancy(ctx context.Context, req repomodel.GetTemplRecomendedFrequancyRequest) ([]model.TemplRecomendedFrequancy, error)
	GetTemplRecomendedFrequancyByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.TemplRecomendedFrequancy, error)
	UpdateTemplRecomendedFrequancyByID(ctx context.Context, input model.TemplRecomendedFrequancy) (*repomodel.BaseIdResponse, error)
}

type TemplQuestionaryTagsRepository interface {
	CreateTemplQuestionaryTags(ctx context.Context, input model.TemplQuestionaryTags) (*repomodel.BaseIdResponse, error)
	CreateTemplQuestionaryTagsRange(ctx context.Context, inputs []model.TemplQuestionaryTags) error
	DeleteTemplQuestionaryTagsByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetTemplQuestionaryTags(ctx context.Context, req repomodel.GetTemplQuestionaryTagsRequest) ([]repomodel.GetTemplQuestionaryTags, error)
	GetTemplQuestionaryTagsByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetTemplQuestionaryTags, error)
	UpdateTemplQuestionaryTagsByID(ctx context.Context, input model.TemplQuestionaryTags) (*repomodel.BaseIdResponse, error)
}

type TemplTemplateQuestionary interface {
	CreateTemplTemplateQuestionary(ctx context.Context, input model.TemplTemplateQuestionary) (*repomodel.BaseIdResponse, error)
	CreateTemplTemplateQuestionaryRange(ctx context.Context, inputs []model.TemplTemplateQuestionary) error
	DeleteTemplTemplateQuestionaryByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetTemplTemplateQuestionary(ctx context.Context, req repomodel.GetTemplTemplateQuestionaryRequest) ([]repomodel.GetTemplTemplateQuestionary, error)
	GetTemplTemplateQuestionaryByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetTemplTemplateQuestionary, error)
	UpdateTemplTemplateQuestionaryByID(ctx context.Context, input model.TemplTemplateQuestionary) (*repomodel.BaseIdResponse, error)
}

type CommunicationTemplateRepository interface {
	CreateCommunicationTemplate(ctx context.Context, createTemplComms *repomodel.CreateCommunicationTemplate) (*repomodel.CommunicationTemplate, error)
	GetCommunicationTemplateByID(ctx context.Context, id int) (*repomodel.CommunicationTemplate, error)
	GetCommunicationTemplatesByEntityID(ctx context.Context, id int) ([]repomodel.CommunicationTemplate, error)
	UpdateCommunicationTemplateByID(ctx context.Context, updateTemplComms *repomodel.UpdateCommunicationTemplate) (*repomodel.CommunicationTemplate, error)
	DeleteCommunicationTemplateByID(ctx context.Context, id int) error
}

var (
	ErrCommunicationTemplateNotFound = errors.New("communication template records not found")
)
