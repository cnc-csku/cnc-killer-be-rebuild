package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ActionDatabase struct {
	config *config.Config
	db     *sqlx.DB
}

func NewActionDatabase(db *sqlx.DB, cfg *config.Config) repositories.ActionRepository {
	return &ActionDatabase{
		config: cfg,
		db:     db,
	}
}

// Implementation of every methods in `action_repository.go`
func (a *ActionDatabase) AddAction(ctx context.Context, action *models.Action) error {
	if action == nil {
		return errors.New("action can't be nil")
	}

	// Generate a new UUID for action_id
	action.ActionID = uuid.New().String()

	fmt.Println(action.ActionID, action.ActionDetail, action.ActionCondition)

	query := `INSERT INTO actions (action_id, action_detail, action_condition) VALUES ($1, $2, $3)`
	if _, err := a.db.ExecContext(ctx, query, action.ActionID, action.ActionDetail, action.ActionCondition); err != nil {
		return err
	}

	return nil
}
