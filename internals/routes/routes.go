package routes

import "github.com/gofiber/fiber/v2"

func SetupSkinsRoutes(router fiber.Router) {
	skins := router.Group("skins")

	skins.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome to my cs2 skins app.")
	})
}
