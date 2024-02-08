package templateservice

import (
	"github.com/engagerocketco/templates-api-svc/internal/repository"
	"github.com/engagerocketco/templates-api-svc/internal/service/natsservice"
	"go.uber.org/zap"
)

type Service interface {
	TemplTemplateQuestionary
	CommunicationTemplateService
	SurveyTags
	TemplRecomendedFrequancy
	TemplQuestionaryTags
}

type service struct {
	repo        repository.Repository
	natsService natsservice.Service
	logger      *zap.Logger
}

func New(repo repository.Repository, natsService natsservice.Service, logger *zap.Logger) Service {
	return &service{
		repo:        repo,
		natsService: natsService,
		logger:      logger,
	}
}
