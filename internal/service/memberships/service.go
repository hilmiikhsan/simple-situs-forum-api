package memberships

import (
	"context"

	"github.com/hilmiikhsan/situs-forum/internal/configs"
	"github.com/hilmiikhsan/situs-forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUseByEmailOrUsername(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
