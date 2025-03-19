package bank

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// GetBanks retrieves all banks
func GetBanks(c *fiber.Ctx) error {
	banks, err := bankRepo.FindAllWithUserAccounts(repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	return response.Send(c, types.ResponseOption{Data: banks})
}

// GetBank retrieves a specific bank by ID
func GetBank(c *fiber.Ctx) error {
	id := c.Params("id")
	bank, err := bankRepo.FindWithUserAccounts(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if bank == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Bank not found"),
			Status: fiber.StatusNotFound,
		})
	}
	return response.Send(c, types.ResponseOption{Data: bank})
}

// GetBankAccounts retrieves all accounts for a specific bank
func GetBankAccounts(c *fiber.Ctx) error {
	id := c.Params("id")
	bank, err := bankRepo.FindWithUserAccounts(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if bank == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Bank not found"),
			Status: fiber.StatusNotFound,
		})
	}
	return response.Send(c, types.ResponseOption{Data: bank.UserAccounts})
}
