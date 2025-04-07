package repositories

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, userID string) (*models.User, error)
	Login(ctx context.Context, passwd string) (*models.User, error)
}
