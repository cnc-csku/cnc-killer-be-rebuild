package postgres

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/jmoiron/sqlx"
)

type ActionDatabase struct {
	db *sqlx.DB
}

func NewActionDatabase(db *sqlx.DB) repositories.ActionRepository {
	return &ActionDatabase{
		db: db,
	}
}

// Implementation of every methods in `action_repository.go`
func (a *ActionDatabase) AddAction(ctx context.Context, action *models.Action) error {
	if action == nil {
		return exceptions.ErrInvalidAction
	}

	query := `INSERT INTO actions (action_id, action_detail, action_condition) VALUES ($1, $2, $3)`
	if _, err := a.db.ExecContext(ctx, query, action.ID, action.Detail, action.Condition); err != nil {
		return err
	}

	return nil
}
