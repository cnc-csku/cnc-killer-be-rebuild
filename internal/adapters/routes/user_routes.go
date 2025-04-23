package routes

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *handlers.Handler) {
	api := app.Group("/user")
	api.Post("/role", handler.UserHandler.GetRole)
}
