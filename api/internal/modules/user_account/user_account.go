package user_account

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/auth"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// GetUserAccounts retrieves all accounts for the current user
func GetUserAccounts(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	accounts, err := accountRepo.FindByUserID(currentUser.ID, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: accounts})
}

// GetUserAccount retrieves a specific account for the current user
func GetUserAccount(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

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

	// Check if the account belongs to the current user
	if account.UserID != currentUser.ID {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusForbidden, "You don't have access to this account"),
			Status: fiber.StatusForbidden,
		})
	}

	return response.Send(c, types.ResponseOption{Data: account})
}

// CreateUserAccount creates a new account for the current user
func CreateUserAccount(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	body := new(struct {
		BankID        *string `json:"bankId"`
		WalletID      *string `json:"walletId"`
		AccountNumber *string `json:"accountNumber"`
		PhoneNumber   *string `json:"phoneNumber"`
		Name          string  `json:"name" validate:"required"`
	})
	if err := c.BodyParser(body); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	newAccount := &models.UserAccount{
		BankID:        body.BankID,
		WalletID:      body.WalletID,
		UserID:        currentUser.ID,
		AccountNumber: body.AccountNumber,
		PhoneNumber:   body.PhoneNumber,
		Name:          body.Name,
		Balance:       0,
		IsActive:      true,
	}

	if err := accountRepo.Create(newAccount, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Data:   newAccount,
		Status: fiber.StatusCreated,
	})
}

// UpdateUserAccount updates an existing account for the current user
func UpdateUserAccount(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

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

	// Check if the account belongs to the current user
	if existingAccount.UserID != currentUser.ID {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusForbidden, "You don't have access to this account"),
			Status: fiber.StatusForbidden,
		})
	}

	body := new(struct {
		BankID        *string `json:"bankId"`
		WalletID      *string `json:"walletId"`
		AccountNumber *string `json:"accountNumber"`
		PhoneNumber   *string `json:"phoneNumber"`
		Name          string  `json:"name" validate:"required"`
		IsActive      *bool   `json:"isActive"`
	})
	if err := c.BodyParser(body); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	existingAccount.BankID = body.BankID
	existingAccount.WalletID = body.WalletID
	existingAccount.AccountNumber = body.AccountNumber
	existingAccount.PhoneNumber = body.PhoneNumber
	existingAccount.Name = body.Name
	if body.IsActive != nil {
		existingAccount.IsActive = *body.IsActive
	}

	if err := accountRepo.Update(existingAccount, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: existingAccount})
}

// DeleteUserAccount deletes an account for the current user
func DeleteUserAccount(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

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

	// Check if the account belongs to the current user
	if existingAccount.UserID != currentUser.ID {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusForbidden, "You don't have access to this account"),
			Status: fiber.StatusForbidden,
		})
	}

	if err := accountRepo.Delete(id, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Status: fiber.StatusNoContent,
	})
}

// GetAccountBalance retrieves the current balance for an account
func GetAccountBalance(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

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

	// Check if the account belongs to the current user
	if existingAccount.UserID != currentUser.ID {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusForbidden, "You don't have access to this account"),
			Status: fiber.StatusForbidden,
		})
	}

	return response.Send(c, types.ResponseOption{Data: existingAccount.Balance})
}

// GetAccountLedgers retrieves all ledger entries for an account
func GetAccountLedgers(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

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

	// Check if the account belongs to the current user
	if existingAccount.UserID != currentUser.ID {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusForbidden, "You don't have access to this account"),
			Status: fiber.StatusForbidden,
		})
	}

	ledgers, err := ledgerRepo.FindByAccountID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: ledgers})
}
