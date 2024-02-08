package natsservice

import (
	"context"

	"github.com/engagerocketco/go-common/ns"
)

type Permissioner interface {
	GetAccountInfo(ctx context.Context, email string) (*ns.AccountInfo, error)
}

func (s *service) GetAccountInfo(ctx context.Context, email string) (*ns.AccountInfo, error) {
	return s.permissionClient.GetAccountInfo(ctx, &ns.AccountInfoRequest{Email: email})
}
