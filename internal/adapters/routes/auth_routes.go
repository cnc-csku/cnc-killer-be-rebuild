package routes

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/rest"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, handler *rest.Handler) {
	api := app.Group("/auth")
	api.Get("/google", handler.GoogleAuthHandler.GoogleLogin)
	api.Get("/google/callback", handler.GoogleAuthHandler.GoogleCallback)
}
