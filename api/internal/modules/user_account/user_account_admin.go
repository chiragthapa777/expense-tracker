package user_account

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// GetAllAccounts retrieves all accounts (admin only)
func GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := accountRepo.Find(repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: accounts})
}

// GetAccount retrieves a specific account (admin only)
func GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	account, err := accountRepo.FindWithAssociations(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if account == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Account not found"),
			Status: fiber.StatusNotFound,
		})
	}

	return response.Send(c, types.ResponseOption{Data: account})
}

// CreateAccount creates a new account (admin only)
func CreateAccount(c *fiber.Ctx) error {
	body := new(struct {
		UserID        string  `json:"userId" validate:"required"`
		BankID        *string `json:"bankId"`
		WalletID      *string `json:"walletId"`
		AccountNumber *string `json:"accountNumber"`
		PhoneNumber   *string `json:"phoneNumber"`
		Name          string  `json:"name" validate:"required"`
		Balance       float64 `json:"balance"`
		IsActive      bool    `json:"isActive"`
	})
	if err := c.BodyParser(body); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	newAccount := &models.UserAccount{
		UserID:        body.UserID,
		BankID:        body.BankID,
		WalletID:      body.WalletID,
		AccountNumber: body.AccountNumber,
		PhoneNumber:   body.PhoneNumber,
		Name:          body.Name,
		Balance:       body.Balance,
		IsActive:      body.IsActive,
	}

	if err := accountRepo.Create(newAccount, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Data:   newAccount,
		Status: fiber.StatusCreated,
	})
}

// UpdateAccount updates an existing account (admin only)
func UpdateAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	existingAccount, err := accountRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if existingAccount == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Account not found"),
			Status: fiber.StatusNotFound,
		})
	}

	body := new(struct {
		UserID        *string  `json:"userId"`
		BankID        *string  `json:"bankId"`
		WalletID      *string  `json:"walletId"`
		AccountNumber *string  `json:"accountNumber"`
		PhoneNumber   *string  `json:"phoneNumber"`
		Name          *string  `json:"name"`
		Balance       *float64 `json:"balance"`
		IsActive      *bool    `json:"isActive"`
	})
	if err := c.BodyParser(body); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	if body.UserID != nil {
		existingAccount.UserID = *body.UserID
	}
	if body.BankID != nil {
		existingAccount.BankID = body.BankID
	}
	if body.WalletID != nil {
		existingAccount.WalletID = body.WalletID
	}
	if body.AccountNumber != nil {
		existingAccount.AccountNumber = body.AccountNumber
	}
	if body.PhoneNumber != nil {
		existingAccount.PhoneNumber = body.PhoneNumber
	}
	if body.Name != nil {
		existingAccount.Name = *body.Name
	}
	if body.Balance != nil {
		existingAccount.Balance = *body.Balance
	}
	if body.IsActive != nil {
		existingAccount.IsActive = *body.IsActive
	}

	if err := accountRepo.Update(existingAccount, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: existingAccount})
}

// DeleteAccount deletes an account (admin only)
func DeleteAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	existingAccount, err := accountRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if existingAccount == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Account not found"),
			Status: fiber.StatusNotFound,
		})
	}

	if err := accountRepo.Delete(id, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Status: fiber.StatusNoContent,
	})
}

// GetUserAccountsAdmin retrieves all accounts for a specific user (admin only)
func GetUserAccountsAdmin(c *fiber.Ctx) error {
	userID := c.Params("userId")
	accounts, err := accountRepo.FindByUserID(userID, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: accounts})
}
