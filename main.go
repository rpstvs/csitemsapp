package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/router"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":8080")

}
