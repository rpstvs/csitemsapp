package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func HomePage(c *fiber.Ctx, db *database.Queries) error {
	items, _ := db.GetBestItems(c.Context())
	var Response []response
	for _, item := range items {
		Response = append(Response, response{
			Itemname: item.Itemname.String,
			Price:    item.Price,
		})
	}
	return c.JSON(Response)
}

func WorstItems(c *fiber.Ctx, db *database.Queries) error {
	items, _ := db.GetWorstItems(c.Context())
	var Response []response
	for _, item := range items {
		Response = append(Response, response{
			Itemname: item.Itemname.String,
			Price:    item.Price,
		})
	}
	return c.JSON(Response)
}
