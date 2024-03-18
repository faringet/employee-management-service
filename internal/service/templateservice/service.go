package templateservice

import (
	"github.com/engagerocketco/templates-api-svc/internal/repository"
	"go.uber.org/zap"
)

type Service interface {
	Attributes
	EmployeeOptionAttributes
	Employees
	SmProject
	SmProjectType
	SmAttributeTriggers
	SmSurvey
	SmSurveyStatus
}

type service struct {
	repo   repository.Repository
	logger *zap.Logger
}

func New(repo repository.Repository, logger *zap.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}
