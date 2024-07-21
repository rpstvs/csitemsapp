package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
	"github.com/rpstvs/csitemsapp/internals/sql"
	"github.com/rpstvs/csitemsapp/router"
)

type DB struct {
	Db database.Queries
}

func main() {

	db := sql.CreateDBInstance()
	app := fiber.New()

	router.SetupRoutes(app, db)

	app.Listen(":8080")

}
