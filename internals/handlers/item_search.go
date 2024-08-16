package handlers

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func ItemSearch(c *fiber.Ctx, db *database.Queries) error {
	query := c.Query("item")

	if query == "" {
		return c.JSON([]string{}) // Return an empty list if no query is provided
	}

	itemSearched, err := url.QueryUnescape(query)

	if err != nil {
		log.Print(err)
	}
	fmt.Println(itemSearched)

	item, err := db.GetItemInfo(c.Context(), itemSearched)

	if err != nil {
		return err
	}

	Price, _ := db.GetLatestPrice(c.Context(), item.ID)

	return c.JSON(response{
		Itemname:   item.Itemname,
		Price:      Price.Price,
		ImageUrl:   item.Imageurl,
		DayChange:  item.Daychange,
		WeekChange: item.Weekchange,
	})
}
