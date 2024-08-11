package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func HomePage(c *fiber.Ctx, db *database.Queries) error {
	items, err := db.GetBestItems(c.Context())
	if err != nil {
		log.Print(err)
	}

	var Response []response
	for _, item := range items {
		price, _ := db.GetLatestPrice(c.Context(), item.ID)
		Response = append(Response, response{
			Itemname:  item.Itemname,
			Price:     price.Price,
			DayChange: item.Daychange,
			ImageUrl:  item.Imageurl,
		})
	}

	return c.JSON(Response)
}

func WorstItems(c *fiber.Ctx, db *database.Queries) error {
	items, _ := db.GetWorstItems(c.Context())
	var Response []response
	for _, item := range items {
		price, _ := db.GetLatestPrice(c.Context(), item.ID)
		Response = append(Response, response{
			Itemname:  item.Itemname,
			Price:     price.Price,
			DayChange: item.Daychange,
			ImageUrl:  item.Imageurl,
		})
	}
	return c.JSON(Response)
}
