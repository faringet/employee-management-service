package natsservice

import (
	"context"

	"github.com/engagerocketco/go-common/ns"
)

type Entity interface {
	GetEntityByID(ctx context.Context, id int) (*ns.GetEntityResponse, error)
}

func (s *service) GetEntityByID(ctx context.Context, id int) (*ns.GetEntityResponse, error) {
	return s.entityClient.GetEntityByID(ctx, &ns.GetEntityRequest{ID: id})
}
