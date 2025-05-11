package repositories

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
)

type PlayerRepository interface {
	AddPlayer(ctx context.Context, player *models.Player) error
	GetPlayerByID(ctx context.Context, playerID string) (*models.Player, error)
	UpdatePlayerIsAlive(ctx context.Context, playerID string, isAlive bool) error
	UpdatePlayerActionID(ctx context.Context, playerID string, victimID string) error
	UpdatePlayerVictimID(ctx context.Context, playerID string, victimID string) error
	UpdatePlayerScore(ctx context.Context, playerID string) error
}
