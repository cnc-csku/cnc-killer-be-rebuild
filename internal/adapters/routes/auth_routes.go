package routes

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, handler *handlers.Handler) {
	api := app.Group("/auth")
	api.Get("/google", handler.GoogleAuthHandler.GoogleLogin)
	api.Get("/google/callback", handler.GoogleAuthHandler.GoogleCallback)
	api.Post("/refresh", handler.GoogleAuthHandler.GetRefreshToken)
}
