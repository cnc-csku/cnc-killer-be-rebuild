package repositories

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
)

type PlayerRepository interface {
	AddPlayer(ctx context.Context, player *models.Player) error
}
