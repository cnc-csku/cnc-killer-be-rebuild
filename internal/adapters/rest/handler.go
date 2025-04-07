package rest

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/postgres"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	UserHandler UserHandler
}

func InitHandler(db *sqlx.DB) *Handler {
	userRepo := postgres.NewUserDatabase(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	return &Handler{
		UserHandler: userHandler,
	}
}
