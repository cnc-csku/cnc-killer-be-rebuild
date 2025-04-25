package postgres

import (
	"context"

	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/google/uuid"
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

// AddUser implements repositories.UserRepository.
func (u *UserDatabase) AddUser(ctx context.Context, email string) error {
	query := `INSERT INTO users (user_id , email , user_role) VALUES ($1, $2 , $3);`
	userID, err := uuid.NewV7()
	if err != nil {
		return err
	}
	role := models.UserRoles.User
	_, err = u.db.ExecContext(ctx, query, userID, email, role)
	if err != nil {
		return err
	}

	return nil

}

// FindUserByEmail implements repositories.UserRepository.
func (u *UserDatabase) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	var user models.User
	err := u.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserRole implements repositories.UserRepository.
func (u *UserDatabase) UpdateUserRole(ctx context.Context, email string, newRole string) error {
	query := `UPEDATE users SET role = $1 WHERE email = $2`
	result, err := u.db.ExecContext(ctx, query, newRole, email)
	if err != nil {
		return err
	}

	if row, _ := result.RowsAffected(); row == 0 {
		return exceptions.ErrUserNotFound
	}

	return nil
}
