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
	ActionHandler     ActionHandler
}

func InitHandler(db *sqlx.DB, cfg *config.Config, googleCfg *config.GoogleAuthConfig) *Handler {
	userRepo := postgres.NewUserDatabase(db, cfg)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	game := config.NewGame()
	managerRepo := facilities.NewManagerInstance(game)
	managerService := services.NewManagerService(managerRepo)
	managerHandler := NewManagerHandler(managerService)

	actionRepo := postgres.NewActionDatabase(db, cfg)
	actionService := services.NewActionService(actionRepo)
	actionHandler := NewActionHandler(actionService)

	authRepo := facilities.NewGoogleAuthInstance(googleCfg)
	authService := services.NewAuthService(authRepo, userRepo)
	googleAuthHandler := NewGoogleAuthHandler(authService)
	return &Handler{
		UserHandler:       userHandler,
		GoogleAuthHandler: googleAuthHandler,
		ManagerHandler:    managerHandler,
		ActionHandler:     actionHandler,
	}

}
