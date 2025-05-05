package services

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/google/uuid"
)

type ActionService interface {
	// FindActionByID
	AddAction(ctx context.Context, actionDetail string, actionCondition string) error
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
func (a *actionServiceImpl) AddAction(ctx context.Context, actionDetail string, actionCondition string) error {
	// Generate a new UUID Version 7 for action_id
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	var action = models.Action{
		ID:        uuid.String(),
		Detail:    actionDetail,
		Condition: actionCondition,
	}

	if err := a.repo.AddAction(ctx, &action); err != nil {
		return err
	}

	return nil
}
