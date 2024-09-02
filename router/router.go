package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rpstvs/csitemsapp/internals/auth"
	"github.com/rpstvs/csitemsapp/internals/database"
	"github.com/rpstvs/csitemsapp/internals/handlers"
)

func SetupRoutes(app *fiber.App, db *database.Queries) {
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.GetBestItemsDay(c, db)
	})

	api := app.Group("/api", logger.New())

	skins := api.Group("/skins")

	skins.Get("/", func(c *fiber.Ctx) error {
		return handlers.GetSkins(c, db)
	})

	skins.Get("/best", func(c *fiber.Ctx) error {
		return handlers.GetBestItemsDay(c, db)
	})

	skins.Get("/best/week", func(c *fiber.Ctx) error {
		return handlers.GetBestItemsWeek(c, db)
	})

	skins.Get("/worst", func(c *fiber.Ctx) error {
		return handlers.WorstItems(c, db)
	})

	skins.Post("/history", func(c *fiber.Ctx) error {
		return handlers.GetPriceHistory(c, db)
	})

	skins.Get("/search", func(c *fiber.Ctx) error {
		return handlers.ItemSearch(c, db)
	})

	skins.Get("/CallbackLogin", func(c *fiber.Ctx) error {
		return handlers.CallbackLogin(c)
	})
	skins.Get("/profile", auth.MiddlewareCookie("SkinsApp"), func(c *fiber.Ctx) error {
		return handlers.Profile(c)
	})

	skins.Get("searchitem", func(c *fiber.Ctx) error {
		return handlers.Suggestions(c, db)
	})

}
