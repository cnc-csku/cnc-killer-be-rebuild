package rest

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/core/services"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/postgres"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/manager"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	UserHandler    UserHandler
	ManagerHandler manager.GameHandler
}

func InitHandler(db *sqlx.DB) *Handler {
	userRepo := postgres.NewUserDatabase(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	managerService := manager.NewGame()
	managerHandler := manager.NewGameHandler(managerService)
	return &Handler{
		UserHandler:    userHandler,
		ManagerHandler: managerHandler,
	}
}
