package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

type UserDatabase struct {
	config *config.Config
	db     *sqlx.DB
}

func NewUserDatabase(db *sqlx.DB, cfg *config.Config) repositories.UserRepository {
	return &UserDatabase{
		db:     db,
		config: cfg,
	}
}

// AddUser implements repositories.UserRepository.
func (u *UserDatabase) AddUser(ctx context.Context, email string) (*models.User, error) {
	query := `INSERT INTO users (email , user_role) VALUES ($1, $2) RETURNING  *`
	role := models.UserRoles.User
	var user models.User
	err := u.db.GetContext(ctx, &user, query, email, role)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

// FindUserByEmail implements repositories.UserRepository.
func (u *UserDatabase) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	var user models.User
	err := u.db.GetContext(ctx, &user, query, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
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

// GenerateJWT implements repositories.UserRepository.
func (u *UserDatabase) GenerateAccessToken(user *models.User) (string, error) {
	duration, err := time.ParseDuration(u.config.JWT.AccessExp)
	if err != nil {
		return "", err
	}

	expire := time.Now().Add(duration).Unix()
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   expire,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(u.config.JWT.AccessExp))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}
