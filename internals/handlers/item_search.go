package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func ItemSearch(c *fiber.Ctx, db *database.Queries) error {
	req := struct {
		Itemname string `json:"itemname"`
	}{}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	itemId, err := db.GetItemByName(c.Context(), req.Itemname)

	if err != nil {
		return err
	}

	Price, _ := db.GetLatestPrice(c.Context(), itemId)

	return c.JSON(response{
		Itemname: req.Itemname,
		Price:    Price.Price,
	})
}
