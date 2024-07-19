package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
	"github.com/rpstvs/csitemsapp/internals/sql"
	"github.com/rpstvs/csitemsapp/router"
)

var DB *database.Queries

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	sql.CreateDB()
	app.Listen(":8080")

}
