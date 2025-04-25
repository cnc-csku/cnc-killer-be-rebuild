package repositories

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, userID string) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	AddUser(ctx context.Context, email string) error
	UpdateUserRole(ctx context.Context, email string, newRole string) error
}
