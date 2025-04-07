package postgres

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/jmoiron/sqlx"
)

type UserDatabase struct {
	db *sqlx.DB
}

func NewUserDatabase(db *sqlx.DB) repositories.UserRepository {
	return &UserDatabase{
		db: db,
	}
}

// GetRole implements repositories.UserRepository.
func (u *UserDatabase) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	query := `SELECT * FROM users WHERE user_id = $1`
	var user models.User
	err := u.db.GetContext(ctx, &user, query, userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Login implements repositories.UserRepository.
func (u *UserDatabase) Login(ctx context.Context, passwd string) (*models.User, error) {
	panic("unimplemented")
}
