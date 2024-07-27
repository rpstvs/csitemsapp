package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareCookie(cookieValue string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cookie := c.Cookies(cookieValue)

		err := ValidateToken(cookie)

		if err != nil {
			log.Print("Cookie isnt valid")
			c.Redirect("/")
		}

		return c.Next()
	}
}
