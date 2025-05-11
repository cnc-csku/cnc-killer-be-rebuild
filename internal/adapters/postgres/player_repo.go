package postgres

import (
	"context"
	"database/sql"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
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

func (p *PlayerDatabase) GetPlayerByID(ctx context.Context, playerID string) (*models.Player, error) {
	var player models.Player
	query := `SELECT * FROM players WHERE player_id = $1`

	if err := p.db.GetContext(ctx, &player, query, playerID); err != nil {
		if err == sql.ErrNoRows {
			return nil, exceptions.ErrPlayerNotFound
		}
		return nil, err
	}

	return &player, nil
}

func (p *PlayerDatabase) UpdatePlayerIsAlive(ctx context.Context, playerID string, isAlive bool) error {
	query := `UPDATE players SET is_alive = $1 WHERE player_id = $2`
	if _, err := p.db.ExecContext(ctx, query, isAlive, playerID); err != nil {
		return err
	}
	return nil
}

func (p *PlayerDatabase) UpdatePlayerActionID(ctx context.Context, playerID string, victimID string) error {
	victim, err := p.GetPlayerByID(ctx, victimID)
	if err != nil {
		return err
	}

	// Update player ActionID // TODO: What if ActionID is nil?
	if _, err := p.db.ExecContext(ctx, `UPDATE players SET action_id = $1 WHERE player_id = $2`, victim.ActionID, playerID); err != nil {
		return err
	}
	// Update victim ActionID
	if _, err := p.db.ExecContext(ctx, `UPDATE players SET action_id = NULL WHERE player_id = $1`, victimID); err != nil {
		return err
	}

	return nil
}

func (p *PlayerDatabase) UpdatePlayerVictimID(ctx context.Context, playerID string, victimID string) error {
	victim, err := p.GetPlayerByID(ctx, victimID)
	if err != nil {
		return err
	}

	// Update player VictimID // TODO: What if VictimID is nil?
	if _, err := p.db.ExecContext(ctx, `UPDATE players SET victim_id = $1 WHERE player_id = $2`, victim.VictimID, playerID); err != nil {
		return err
	}
	// Update victim VictimID
	if _, err := p.db.ExecContext(ctx, `UPDATE players SET victim_id = NULL WHERE player_id = $1`, victimID); err != nil {
		return err
	}

	return nil
}

func (p *PlayerDatabase) UpdatePlayerScore(ctx context.Context, playerID string) error {
	// TODO: Implements add score logics
	return nil
}
