package ledger

import (
	"time"

	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// AdminGetAllLedgers retrieves all ledgers (admin only)
func AdminGetAllLedgers(c *fiber.Ctx) error {
	ledgers, err := ledgerRepo.Find(repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: ledgers})
}

// AdminGetLedger retrieves a specific ledger (admin only)
func AdminGetLedger(c *fiber.Ctx) error {
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

	return response.Send(c, types.ResponseOption{Data: ledger})
}

// AdminGetLedgersByUser retrieves all ledgers for a specific user (admin only)
func AdminGetLedgersByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	ledgers, err := ledgerRepo.FindByUserID(userID, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: ledgers})
}

// AdminGetLedgersByAccount retrieves all ledgers for a specific account (admin only)
func AdminGetLedgersByAccount(c *fiber.Ctx) error {
	accountID := c.Params("accountId")
	ledgers, err := ledgerRepo.FindByAccountID(accountID, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: ledgers})
}

// AdminGetLedgersByDateRange retrieves all ledgers within a date range (admin only)
func AdminGetLedgersByDateRange(c *fiber.Ctx) error {
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

	return response.Send(c, types.ResponseOption{Data: ledgers})
}
