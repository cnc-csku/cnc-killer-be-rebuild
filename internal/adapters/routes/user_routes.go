package routes

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/rest"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *rest.Handler) {
	api := app.Group("/user")
	api.Post("/role", handler.UserHandler.GetRole)
}
