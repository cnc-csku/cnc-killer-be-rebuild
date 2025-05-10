package routes

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/handlers"
	"github.com/gofiber/fiber/v2"
)

func PlayerRoutes(app *fiber.App, handler *handlers.Handler) {
	api := app.Group("/player")
	api.Post("/", handler.PlayerHandler.AddPlayer)
}
