package repositories

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	AddUser(ctx context.Context, email string) (*models.User, error)
	UpdateUserRole(ctx context.Context, email string, newRole string) error
	GenerateAccessToken(user *models.User, isAccessToken bool) (string, error)
	GenerateRefreshToken(ctx context.Context, accessToken string) (string, error)
}
