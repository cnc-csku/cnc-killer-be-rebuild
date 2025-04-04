package main

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.NewConfig()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "connected",
		})
	})
	if err := app.Listen(":" + cfg.Port); err != nil {
		panic(err)
	}
}
