package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func Suggestions(c *fiber.Ctx, db *database.Queries) error {
	query := c.Query("q")

	if query == "" {
		return c.JSON([]string{}) // Return an empty list if no query is provided
	}

	items, _ := db.GetItemSuggestion(c.Context(), "%"+query+"%")
	suggestions := []string{}

	for _, item := range items {
		suggestions = append(suggestions, item)
	}

	return c.JSON(suggestions)
}
