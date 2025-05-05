package services

import (
	"context"
	"errors"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
)

type ActionService interface {
	// FindActionByID
	AddAction(ctx context.Context, action *models.Action) error
}

type actionServiceImpl struct {
	repo repositories.ActionRepository
}

func NewActionService(repo repositories.ActionRepository) ActionService {
	return &actionServiceImpl{
		repo: repo,
	}
}

// Implementation of every methods in `UserService`
func (a *actionServiceImpl) AddAction(ctx context.Context, action *models.Action) error {
	if err := a.repo.AddAction(ctx, action); err != nil {
		return errors.New("failed to add an action")
	}
	return nil
}
