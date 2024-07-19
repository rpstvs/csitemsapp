package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome to my cs2 skins app.")
	})

	api := app.Group("/api", logger.New())

	skins := api.Group("/skins")

	skins.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("this is the skins endpoint")
	})

}
