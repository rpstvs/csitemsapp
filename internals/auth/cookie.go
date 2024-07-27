package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCookie(tokenstring string) fiber.Cookie {
	return fiber.Cookie{
		Name:     "SkinsApp",
		Value:    tokenstring,
		Expires:  time.Now().UTC().Add(6 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	}
}
