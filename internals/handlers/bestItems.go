package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func GetBestITems(c *fiber.Ctx, db *database.Queries) error {
	items, _ := db.GetBestItems(c.Context())
	var name string
	var price float64
	resp := []response{}

	for _, item := range items {
		name, _ = db.GetNameById(c.Context(), item.ItemID)
		price = item.Price
		resp = append(resp, response{
			Itemname: name,
			Price:    price,
		})
	}

	return c.JSON(resp)
}
