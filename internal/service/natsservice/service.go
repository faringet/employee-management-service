package natsservice

import (
	"github.com/engagerocketco/go-common/ns"
	"go.uber.org/zap"
)

type Service interface {
	Entity
	Permissioner
}

type service struct {
	entityClient     ns.EntityService
	permissionClient ns.PermissionService
	logger           *zap.Logger
}

func NewNatsService(entityClient ns.EntityService, permissionClient ns.PermissionService, logger *zap.Logger) Service {
	return &service{
		entityClient:     entityClient,
		permissionClient: permissionClient,
		logger:           logger.Named("natsservice"),
	}
}
