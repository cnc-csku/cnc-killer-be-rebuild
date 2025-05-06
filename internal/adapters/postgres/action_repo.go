package postgres

import (
	"context"
	"database/sql"

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

func (a *ActionDatabase) FindActionByID(ctx context.Context, actionID string) (*models.Action, error) {
	var action models.Action
	query := `SELECT * FROM actions WHERE action_id = $1`

	if err := a.db.GetContext(ctx, &action, query, actionID); err != nil {
		if err == sql.ErrNoRows {
			return nil, exceptions.ErrActionNotFound
		}
		return nil, err
	}

	return &action, nil
}
