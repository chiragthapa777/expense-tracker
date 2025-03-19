package bank

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// CreateBank creates a new bank (admin only)
func CreateBank(c *fiber.Ctx) error {
	body := new(struct {
		Name string `json:"name" validate:"required"`
	})
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	newBank := &models.Bank{
		Name: body.Name,
	}

	if err := bankRepo.Create(newBank, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Data:   newBank,
		Status: fiber.StatusCreated,
	})
}

// UpdateBank updates an existing bank (admin only)
func UpdateBank(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(struct {
		Name string `json:"name" validate:"required"`
	})
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	existingBank, err := bankRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if existingBank == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Bank not found"),
			Status: fiber.StatusNotFound,
		})
	}

	existingBank.Name = body.Name
	if err := bankRepo.Update(existingBank, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: existingBank})
}

// DeleteBank deletes a bank (admin only)
func DeleteBank(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := bankRepo.Delete(id, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Status: fiber.StatusNoContent,
	})
}
