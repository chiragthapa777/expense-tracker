package ledger

import (
	"time"

	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/auth"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// GetLedgers retrieves all ledgers for the current user
func GetLedgers(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	ledgers, err := ledgerRepo.FindByUserID(currentUser.ID, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: ledgers})
}

// GetLedger retrieves a specific ledger for the current user
func GetLedger(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	id := c.Params("id")
	ledger, err := ledgerRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if ledger == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Ledger not found"),
			Status: fiber.StatusNotFound,
		})
	}

	// Check if the ledger belongs to the current user
	if ledger.UserID != currentUser.ID {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusForbidden, "You don't have access to this ledger"),
			Status: fiber.StatusForbidden,
		})
	}

	return response.Send(c, types.ResponseOption{Data: ledger})
}

// CreateLedger creates a new ledger for the current user
func CreateLedger(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	body := new(struct {
		Debit         float64 `json:"debit" validate:"required"`
		Credit        float64 `json:"credit" validate:"required"`
		AccountID     string  `json:"accountId" validate:"required"`
		TransactionID string  `json:"transactionId"`
		Description   string  `json:"description" validate:"required"`
		Date          string  `json:"date" validate:"required"`
	})
	if err := c.BodyParser(body); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	date, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD"),
			Status: fiber.StatusBadRequest,
		})
	}

	accountID := body.AccountID
	transactionID := body.TransactionID
	newLedger := &models.Ledger{
		Debit:         body.Debit,
		Credit:        body.Credit,
		AccountID:     &accountID,
		TransactionID: &transactionID,
		UserID:        currentUser.ID,
		Description:   body.Description,
		Date:          date,
	}

	if err := ledgerRepo.Create(newLedger, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Data:   newLedger,
		Status: fiber.StatusCreated,
	})
}

// UpdateLedger updates an existing ledger for the current user
func UpdateLedger(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	id := c.Params("id")
	existingLedger, err := ledgerRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if existingLedger == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Ledger not found"),
			Status: fiber.StatusNotFound,
		})
	}

	// Check if the ledger belongs to the current user
	if existingLedger.UserID != currentUser.ID {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusForbidden, "You don't have access to this ledger"),
			Status: fiber.StatusForbidden,
		})
	}

	body := new(struct {
		Debit         *float64 `json:"debit"`
		Credit        *float64 `json:"credit"`
		AccountID     *string  `json:"accountId"`
		TransactionID *string  `json:"transactionId"`
		Description   *string  `json:"description"`
		Date          *string  `json:"date"`
	})
	if err := c.BodyParser(body); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	if body.Debit != nil {
		existingLedger.Debit = *body.Debit
	}
	if body.Credit != nil {
		existingLedger.Credit = *body.Credit
	}
	if body.AccountID != nil {
		accountID := *body.AccountID
		existingLedger.AccountID = &accountID
	}
	if body.TransactionID != nil {
		transactionID := *body.TransactionID
		existingLedger.TransactionID = &transactionID
	}
	if body.Description != nil {
		existingLedger.Description = *body.Description
	}
	if body.Date != nil {
		date, err := time.Parse("2006-01-02", *body.Date)
		if err != nil {
			return response.SendError(c, types.ErrorResponseOption{
				Error:  fiber.NewError(fiber.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD"),
				Status: fiber.StatusBadRequest,
			})
		}
		existingLedger.Date = date
	}

	if err := ledgerRepo.Update(existingLedger, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: existingLedger})
}

// DeleteLedger deletes a ledger for the current user
func DeleteLedger(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	id := c.Params("id")
	existingLedger, err := ledgerRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if existingLedger == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusNotFound, "Ledger not found"),
			Status: fiber.StatusNotFound,
		})
	}

	// Check if the ledger belongs to the current user
	if existingLedger.UserID != currentUser.ID {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusForbidden, "You don't have access to this ledger"),
			Status: fiber.StatusForbidden,
		})
	}

	if err := ledgerRepo.Delete(id, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Status: fiber.StatusNoContent,
	})
}

// GetLedgersByAccount retrieves all ledgers for a specific account
func GetLedgersByAccount(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	accountID := c.Params("accountId")
	ledgers, err := ledgerRepo.FindByAccountID(accountID, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	// Filter ledgers to only show those belonging to the current user
	var userLedgers []*models.Ledger
	for _, ledger := range ledgers {
		if ledger.UserID == currentUser.ID {
			userLedgers = append(userLedgers, &ledger)
		}
	}

	return response.Send(c, types.ResponseOption{Data: userLedgers})
}

// GetLedgersByDateRange retrieves all ledgers within a date range for the current user
func GetLedgersByDateRange(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusBadRequest, "Invalid start date format. Use YYYY-MM-DD"),
			Status: fiber.StatusBadRequest,
		})
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  fiber.NewError(fiber.StatusBadRequest, "Invalid end date format. Use YYYY-MM-DD"),
			Status: fiber.StatusBadRequest,
		})
	}

	ledgers, err := ledgerRepo.FindByDateRange(start, end, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	// Filter ledgers to only show those belonging to the current user
	var userLedgers []*models.Ledger
	for _, ledger := range ledgers {
		if ledger.UserID == currentUser.ID {
			userLedgers = append(userLedgers, &ledger)
		}
	}

	return response.Send(c, types.ResponseOption{Data: userLedgers})
}
