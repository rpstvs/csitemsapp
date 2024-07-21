package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rpstvs/csitemsapp/internals/database"
	"github.com/rpstvs/csitemsapp/internals/handlers"
)

func SetupRoutes(app *fiber.App, db *database.Queries) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome to my cs2 skins app.")
	})

	api := app.Group("/api", logger.New())

	skins := api.Group("/skins")

	skins.Get("/", func(c *fiber.Ctx) error {
		return handlers.GetSkins(c, db)
	})

	skins.Get("/Best", func(c *fiber.Ctx) error {
		return handlers.GetBestITems(c, db)
	})

}
