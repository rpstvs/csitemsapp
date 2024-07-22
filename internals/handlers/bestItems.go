package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func GetItems(c *fiber.Ctx, db *database.Queries) error {
	items, _ := db.GetItemsRecord(c.Context())

	resp := []response{}

	for _, item := range items {
		price, _ := strconv.ParseFloat(item.Price.String, 64)
		resp = append(resp, response{
			Itemname: item.Itemname,
			Price:    price,
		})

	}

	return c.JSON(resp)
}
