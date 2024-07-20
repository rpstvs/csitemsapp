package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

type response struct {
	Itemname string  `json:"itemname"`
	Price    float64 `json:"price"`
}

func GetSkins(c *fiber.Ctx, db *database.Queries) error {
	id, err := db.GetItemByName(c.Context(), "Chroma Case")

	if err != nil {
		log.Printf("Item not found")
	}

	price, err := db.GetPricebyID(c.Context(), id)

	resp := &response{
		Itemname: "Chroma Case",
		Price:    price,
	}

	return c.JSON(resp)
}
