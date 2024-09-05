package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func GetItems(c *fiber.Ctx, db *database.Queries) error {
	items, _ := db.GetItemsRecord(c.Context(), 3)

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

func GetBestItemsWeek(c *fiber.Ctx, db *database.Queries) error {
	items, err := db.GetBestItemsWeek(c.Context())
	if err != nil {
		log.Print(err)
	}

	var Response []response
	for _, item := range items {
		price, _ := db.GetLatestPrice(c.Context(), item.ID)
		Response = append(Response, response{
			Itemname:   item.Itemname,
			Price:      price.Price,
			WeekChange: item.Weekchange,
			ImageUrl:   item.Imageurl,
		})
	}

	return c.JSON(Response)
}
