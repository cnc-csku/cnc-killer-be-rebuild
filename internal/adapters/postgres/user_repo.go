package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/exceptions"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/models"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/mitchellh/mapstructure"
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
func (u *UserDatabase) GenerateAccessToken(user *models.User, isAccessToken bool) (string, error) {
	var durationStr string
	if isAccessToken {
		durationStr = u.config.JWT.AccessExp
	} else {
		durationStr = u.config.JWT.RefreshExp
	}
	duration, err := time.ParseDuration(durationStr)
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
	signedToken, err := token.SignedString([]byte(u.config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}

// GenerateRefreshToken implements repositories.UserRepository.
func (u *UserDatabase) GenerateRefreshToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.config.JWT.Secret), nil
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	exp := int64(claims["exp"].(float64))

	if email == "" {
		return "", exceptions.ErrEmailNotFound
	}

	user, err := u.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", exceptions.ErrUserNotFound
	}

	RefreshTokenStr, err := u.GenerateAccessToken(user, false)

	if err != nil {
		return "", err
	}

	query := `UPDATE users SET refresh_token = $1 WHERE email = $2;`
	_, err = u.db.ExecContext(ctx, query, RefreshTokenStr, email)
	if err != nil {
		return "", err
	}

	if exp-time.Now().Unix() < 0 {
		accessToken, err = u.GenerateAccessToken(user, true)
	}

	return accessToken, err

}

// ExactJWT implements repositories.UserRepository.
func (u *UserDatabase) ExactJWT(tokenStr string) (*models.JWTToken, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(u.config.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, exceptions.ErrInvalidJWT
	}
	var jwtToken models.JWTToken
	err = mapstructure.Decode(&claims, &jwtToken)
	if err != nil {
		return nil, err
	}
	return &jwtToken, nil
}

// UpdateUserNickname implements repositories.UserRepository.
func (u *UserDatabase) UpdateUserNickname(ctx context.Context, email string, newNickname string) error {
	log.Println(email, newNickname)
	user, err := u.FindUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user == nil {
		return exceptions.ErrUserNotFound
	}
	query := `UPDATE users SET nickname=$1 WHERE email=$2`
	_, err = u.db.ExecContext(ctx, query, newNickname, email)
	if err != nil {
		return err
	}
	return nil
}
