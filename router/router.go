package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/routes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.SetupSkinsRoutes(api)
}
