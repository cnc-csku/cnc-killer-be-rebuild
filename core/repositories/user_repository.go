package repositories

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	AddUser(ctx context.Context, email string) error
	UpdateUserRole(ctx context.Context, email string, newRole string) error
}
