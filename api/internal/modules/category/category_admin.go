package category

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// CreateCategory creates a new category (admin only)
func CreateCategory(c *fiber.Ctx) error {
	body := new(struct {
		Name        string `json:"name" validate:"required"`
		Icon        string `json:"icon"`
		Color       string `json:"color" validate:"required"`
		Description string `json:"description" validate:"required"`
	})
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	newCategory := &models.Category{
		Name:        body.Name,
		Icon:        body.Icon,
		Color:       body.Color,
		Description: body.Description,
	}

	if err := categoryRepo.Create(newCategory, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Data:   newCategory,
		Status: fiber.StatusCreated,
	})
}

// UpdateCategory updates an existing category (admin only)
func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(struct {
		Name        string `json:"name" validate:"required"`
		Icon        string `json:"icon"`
		Color       string `json:"color" validate:"required"`
		Description string `json:"description" validate:"required"`
	})
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	existingCategory, err := categoryRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if existingCategory == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Category not found"),
			Status: fiber.StatusNotFound,
		})
	}

	existingCategory.Name = body.Name
	existingCategory.Icon = body.Icon
	existingCategory.Color = body.Color
	existingCategory.Description = body.Description

	if err := categoryRepo.Update(existingCategory, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: existingCategory})
}

// DeleteCategory deletes a category (admin only)
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := categoryRepo.Delete(id, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Status: fiber.StatusNoContent,
	})
}
