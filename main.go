package main

import (
	"context"
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/handlers"
	"github.com/cnc-csku/cnc-killer-be-rebuild/internal/adapters/routes"
	"github.com/gofiber/fiber/v2"

	_ "github.com/cnc-csku/cnc-killer-be-rebuild/docs"
)

//@title cnc-killer-api
//@version 1.0
//@description this is for cnc killer backend
//@host localhost:8000
//@BasePath /

func main() {
	cfg := config.NewConfig()
	ctx := context.Background()
	app := fiber.New()

	db := config.ConnectDatabase(cfg, ctx)
	googleCfg := config.NewGoogleConfig(cfg)

	defer db.Close()

	handler := handlers.InitHandler(db, cfg, googleCfg)

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "connected",
		})
	})

	routes.ManagerRoutes(app, handler)
	routes.UserRoutes(app, handler)
	routes.AuthRoute(app, handler)
	routes.ActionRoutes(app, handler)

	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
