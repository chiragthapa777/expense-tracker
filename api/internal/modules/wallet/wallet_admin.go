package wallet

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// CreateWallet creates a new wallet (admin only)
func CreateWallet(c *fiber.Ctx) error {
	body := new(struct {
		Name string `json:"name" validate:"required"`
	})
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	newWallet := &models.Wallet{
		Name: body.Name,
	}

	if err := walletRepo.Create(newWallet, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Data:   newWallet,
		Status: fiber.StatusCreated,
	})
}

// UpdateWallet updates an existing wallet (admin only)
func UpdateWallet(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(struct {
		Name string `json:"name" validate:"required"`
	})
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	existingWallet, err := walletRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if existingWallet == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Wallet not found"),
			Status: fiber.StatusNotFound,
		})
	}

	existingWallet.Name = body.Name
	if err := walletRepo.Update(existingWallet, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: existingWallet})
}

// DeleteWallet deletes a wallet (admin only)
func DeleteWallet(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := walletRepo.Delete(id, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Status: fiber.StatusNoContent,
	})
}
