package rest

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/postgres"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/manager"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	UserHandler       UserHandler
	GoogleAuthHandler GoogleAuthHandler
	ManagerHandler    manager.GameHandler
}

func InitHandler(db *sqlx.DB, googleCfg *config.GoogleAuthConfig) *Handler {
	userRepo := postgres.NewUserDatabase(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	managerService := manager.NewGame()
	managerHandler := manager.NewGameHandler(managerService)

	googleAuthHandler := NewGoogleAuthHandler(googleCfg)
	return &Handler{
		UserHandler:       userHandler,
		GoogleAuthHandler: googleAuthHandler,
		ManagerHandler:    managerHandler,
	}
}
