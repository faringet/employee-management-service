package repository

import (
	"context"
	"errors"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"
	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel/.jet/model"
)

type Repository interface {
	EmployeesRepository
	EmployeeOptionAttributesRepository
	AttributesRepository
	SmProjectRepository
	SmProjectTypeRepository
	SmAttributeTriggersRepository
	SmSurveyRepository
	SmSurveyStatusRepository
	SmSurveyRecepientsRepository
}

var (
	ErrCommunicationTemplateNotFound = errors.New("communication template records not found")
)

type EmployeesRepository interface {
	CreateEmployees(ctx context.Context, input model.Employees) (*repomodel.BaseIdResponse, error)
	CreateEmployeesRange(ctx context.Context, inputs []model.Employees) error
	DeleteEmployeesByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetEmployees(ctx context.Context, req repomodel.GetEmployeesRequest) ([]repomodel.GetEmployees, error)
	GetEmployeesByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetEmployees, error)
	UpdateEmployeesByID(ctx context.Context, input model.Employees) (*repomodel.BaseIdResponse, error)
	CountEmployees(ctx context.Context) (*int, error)
}
type EmployeeOptionAttributesRepository interface {
	CreateEmployeeOptionAttributes(ctx context.Context, input model.EmployeeOptionAttributes) (*repomodel.BaseIdResponse, error)
	CreateEmployeeOptionAttributesRange(ctx context.Context, inputs []model.EmployeeOptionAttributes) error
	DeleteEmployeeOptionAttributesByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetEmployeeOptionAttributes(ctx context.Context, req repomodel.GetEmployeeOptionAttributesRequest) ([]repomodel.GetEmployeeOptionAttributes, error)
	GetEmployeeOptionAttributesByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetEmployeeOptionAttributes, error)
	UpdateEmployeeOptionAttributesByID(ctx context.Context, input model.EmployeeOptionAttributes) (*repomodel.BaseIdResponse, error)
	CountEmployeeOptionAttributes(ctx context.Context) (*int, error)
}
type AttributesRepository interface {
	CreateAttributes(ctx context.Context, input model.Attributes) (*repomodel.BaseIdResponse, error)
	CreateAttributesRange(ctx context.Context, inputs []model.Attributes) error
	DeleteAttributesByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetAttributes(ctx context.Context, req repomodel.GetAttributesRequest) ([]model.Attributes, error)
	GetAttributesByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.Attributes, error)
	UpdateAttributesByID(ctx context.Context, input model.Attributes) (*repomodel.BaseIdResponse, error)
	CountAttributes(ctx context.Context) (*int, error)
}

type SmProjectRepository interface {
	CreateSmProject(ctx context.Context, input model.SmProject) (*repomodel.BaseIdResponse, error)
	CreateSmProjectRange(ctx context.Context, inputs []model.SmProject) error
	DeleteSmProjectByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetSmProject(ctx context.Context, req repomodel.GetSmProjectRequest) ([]repomodel.GetSmProject, error)
	GetSmProjectByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetSmProject, error)
	UpdateSmProjectByID(ctx context.Context, input model.SmProject) (*repomodel.BaseIdResponse, error)
	CountSmProject(ctx context.Context) (*int, error)
}
type SmAttributeTriggersRepository interface {
	CreateSmAttributeTriggers(ctx context.Context, input model.SmAttributeTriggers) (*repomodel.BaseIdResponse, error)
	CreateSmAttributeTriggersRange(ctx context.Context, inputs []model.SmAttributeTriggers) error
	DeleteSmAttributeTriggersByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetSmAttributeTriggers(ctx context.Context, req repomodel.GetSmAttributeTriggersRequest) ([]repomodel.GetSmAttributeTriggers, error)
	GetSmAttributeTriggersByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetSmAttributeTriggers, error)
	UpdateSmAttributeTriggersByID(ctx context.Context, input model.SmAttributeTriggers) (*repomodel.BaseIdResponse, error)
	CountSmAttributeTriggers(ctx context.Context) (*int, error)
}
type SmSurveyRepository interface {
	CreateSmSurvey(ctx context.Context, input model.SmSurvey) (*repomodel.BaseIdResponse, error)
	CreateSmSurveyRange(ctx context.Context, inputs []model.SmSurvey) error
	DeleteSmSurveyByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetSmSurvey(ctx context.Context, req repomodel.GetSmSurveyRequest) ([]repomodel.GetSmSurvey, error)
	GetSmSurveyByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetSmSurvey, error)
	UpdateSmSurveyByID(ctx context.Context, input model.SmSurvey) (*repomodel.BaseIdResponse, error)
	CountSmSurvey(ctx context.Context) (*int, error)
}
type SmProjectTypeRepository interface {
	CreateSmProjectType(ctx context.Context, input model.SmProjectType) (*repomodel.BaseIdResponse, error)
	CreateSmProjectTypeRange(ctx context.Context, inputs []model.SmProjectType) error
	DeleteSmProjectTypeByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetSmProjectType(ctx context.Context, req repomodel.GetSmProjectTypeRequest) ([]model.SmProjectType, error)
	GetSmProjectTypeByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.SmProjectType, error)
	UpdateSmProjectTypeByID(ctx context.Context, input model.SmProjectType) (*repomodel.BaseIdResponse, error)
	CountSmProjectType(ctx context.Context) (*int, error)
}
type SmSurveyStatusRepository interface {
	CreateSmSurveyStatus(ctx context.Context, input model.SmSurveyStatus) (*repomodel.BaseIdResponse, error)
	CreateSmSurveyStatusRange(ctx context.Context, inputs []model.SmSurveyStatus) error
	DeleteSmSurveyStatusByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetSmSurveyStatus(ctx context.Context, req repomodel.GetSmSurveyStatusRequest) ([]model.SmSurveyStatus, error)
	GetSmSurveyStatusByID(ctx context.Context, input repomodel.BaseIdRequest) (*model.SmSurveyStatus, error)
	UpdateSmSurveyStatusByID(ctx context.Context, input model.SmSurveyStatus) (*repomodel.BaseIdResponse, error)
	CountSmSurveyStatus(ctx context.Context) (*int, error)
}

type SmSurveyRecepientsRepository interface {
	CreateSmSurveyRecepients(ctx context.Context, input model.SmSurveyRecepients) (*repomodel.BaseIdResponse, error)
	CreateSmSurveyRecepientsRange(ctx context.Context, inputs []model.SmSurveyRecepients) error
	DeleteSmSurveyRecepientsByID(ctx context.Context, input repomodel.BaseIdRequest) error
	GetSmSurveyRecepients(ctx context.Context, req repomodel.GetSmSurveyRecepientsRequest) ([]repomodel.GetSmSurveyRecepients, error)
	GetSmSurveyRecepientsByID(ctx context.Context, input repomodel.BaseIdRequest) (*repomodel.GetSmSurveyRecepients, error)
	UpdateSmSurveyRecepientsByID(ctx context.Context, input model.SmSurveyRecepients) (*repomodel.BaseIdResponse, error)
	CountSmSurveyRecepients(ctx context.Context) (*int, error)
}
