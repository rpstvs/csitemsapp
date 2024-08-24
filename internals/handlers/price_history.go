package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

type Skin struct {
	Skinanme string  `json:"skinname" form:"skinname"`
	Price    float64 `json:"price" form:"price"`
}

func GetPriceHistory(c *fiber.Ctx, db *database.Queries) error {
	skin := new(Skin)

	if err := c.BodyParser(skin); err != nil {
		return err
	}

	itemHistory, _ := db.GetPriceHistory(c.Context(), "Chroma Case")
	skin.Price = itemHistory.PricesPrice

	return c.JSON(skin)
}
