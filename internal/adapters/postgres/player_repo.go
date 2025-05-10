package postgres

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/jmoiron/sqlx"
)

type PlayerDatabase struct {
	db *sqlx.DB
}

func NewPlayerDatabase(db *sqlx.DB) repositories.PlayerRepository {
	return &PlayerDatabase{
		db: db,
	}
}

func (p *PlayerDatabase) AddPlayer(ctx context.Context, player *models.Player) error {
	query := `INSERT INTO players(player_id, secret_code, victim_id, is_alive, score, action_id, email) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	if _, err := p.db.ExecContext(
		ctx, query, player.ID, player.SecretCode, player.VictimID,
		player.IsAlive, player.Score, player.ActionID, player.Email,
	); err != nil {
		return err
	}
	return nil
}
