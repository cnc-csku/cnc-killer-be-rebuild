package repositories

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
)

type ActionRepository interface {
	AddAction(ctx context.Context, action *models.Action) error
}
