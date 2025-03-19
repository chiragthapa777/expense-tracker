package category

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// GetCategories retrieves all categories
func GetCategories(c *fiber.Ctx) error {
	categories, err := categoryRepo.Find(repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	return response.Send(c, types.ResponseOption{Data: categories})
}

// GetCategory retrieves a specific category by ID
func GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	category, err := categoryRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if category == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Category not found"),
			Status: fiber.StatusNotFound,
		})
	}
	return response.Send(c, types.ResponseOption{Data: category})
}

// GetCategoriesByColor retrieves all categories with a specific color
func GetCategoriesByColor(c *fiber.Ctx) error {
	color := c.Params("color")
	categories, err := categoryRepo.FindByColor(color, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	return response.Send(c, types.ResponseOption{Data: categories})
}
