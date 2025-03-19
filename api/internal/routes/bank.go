package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/bank"
	"github.com/gofiber/fiber/v2"
)

func SetupBankRoutes(app fiber.Router) {
	log := logger.GetLogger()
	bankRoute := app.Group("/bank", middleware.AuthGuard)
	adminBankRoute := app.Group("/admin/bank", middleware.AuthGuardAdminOnly)

	// user routes
	bankRoute.Get("/", bank.GetBanks)
	bankRoute.Get("/:id", bank.GetBank)
	bankRoute.Get("/:id/accounts", bank.GetBankAccounts)

	// admin routes
	adminBankRoute.Post("/", bank.CreateBank)
	adminBankRoute.Put("/:id", bank.UpdateBank)
	adminBankRoute.Delete("/:id", bank.DeleteBank)

	log.Info("Bank routes initialized")
}
