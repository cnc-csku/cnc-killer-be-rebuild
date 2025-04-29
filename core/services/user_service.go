package services

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/responses"
)

type UserService interface {
	GetUserRole(ctx context.Context, email string) (*responses.RoleResponse, error)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}

type userServiceImpl struct {
	repo repositories.UserRepository
}

// GetUserRole implements UserService.
func (u *userServiceImpl) GetUserRole(ctx context.Context, email string) (*responses.RoleResponse, error) {
	user, err := u.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &responses.RoleResponse{
		Role: user.Role,
	}, nil
}
