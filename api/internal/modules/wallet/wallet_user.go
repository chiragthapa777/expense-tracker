package wallet

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// GetWallets retrieves all wallets
func GetWallets(c *fiber.Ctx) error {
	wallets, err := walletRepo.FindAllWithUserAccounts(repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	return response.Send(c, types.ResponseOption{Data: wallets})
}

// GetWallet retrieves a specific wallet by ID
func GetWallet(c *fiber.Ctx) error {
	id := c.Params("id")
	wallet, err := walletRepo.FindWithUserAccounts(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if wallet == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Wallet not found"),
			Status: fiber.StatusNotFound,
		})
	}
	return response.Send(c, types.ResponseOption{Data: wallet})
}

// GetWalletAccounts retrieves all accounts for a specific wallet
func GetWalletAccounts(c *fiber.Ctx) error {
	id := c.Params("id")
	wallet, err := walletRepo.FindWithUserAccounts(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if wallet == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Wallet not found"),
			Status: fiber.StatusNotFound,
		})
	}
	return response.Send(c, types.ResponseOption{Data: wallet.UserAccounts})
}
