package handlers

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/facilities"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/postgres"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	UserHandler       UserHandler
	GoogleAuthHandler GoogleAuthHandler
	ManagerHandler    ManagerHandler
}

func InitHandler(db *sqlx.DB, googleCfg *config.GoogleAuthConfig) *Handler {
	userRepo := postgres.NewUserDatabase(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	game := config.NewGame()
	managerRepo := facilities.NewManagerInstance(game)
	managerService := services.NewManagerService(managerRepo)
	managerHandler := NewManagerHandler(managerService)

	authRepo := facilities.NewGoogleAuthInstance(googleCfg)
	authService := services.NewAuthService(authRepo)
	googleAuthHandler := NewGoogleAuthHandler(authService)
	return &Handler{
		UserHandler:       userHandler,
		GoogleAuthHandler: googleAuthHandler,
		ManagerHandler:    managerHandler,
	}
}
