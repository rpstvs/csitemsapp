package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

type response struct {
	Itemname   string  `json:"itemname"`
	Price      float64 `json:"price"`
	DayChange  float64 `json:"daychange"`
	WeekChange float64 `json:"weekchange"`
	ImageUrl   string  `json:"imageurl"`
}

func GetSkins(c *fiber.Ctx, db *database.Queries) error {
	id, err := db.GetItemByName(c.Context(), "Chroma Case")

	if err != nil {
		log.Printf("Item not found")
	}

	price, _ := db.GetPricebyID(c.Context(), id)

	resp := &response{
		Itemname: "Chroma Case",
		Price:    price,
	}

	return c.JSON(resp)
}
