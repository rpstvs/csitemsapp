package handlers

import "github.com/gofiber/fiber/v2"

func Profile(c *fiber.Ctx) error {
	return c.SendString("Welcome to the app")
}
