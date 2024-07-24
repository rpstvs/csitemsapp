package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func GetPriceHistory(c *fiber.Ctx, db *database.Queries) error {

	itemHistory, _ := db.GetPriceHistory(c.Context(), "Chroma Case")

	return c.JSON(&response{
		Itemname: "Chroma Case",
		Price:    itemHistory,
	})
}
