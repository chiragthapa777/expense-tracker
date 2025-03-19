package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/ledger"
	"github.com/gofiber/fiber/v2"
)

func SetupLedgerRoutes(app fiber.Router) {
	log := logger.GetLogger()
	ledgerRoute := app.Group("/ledger", middleware.AuthGuard)
	adminLedgerRoute := app.Group("/admin/ledger", middleware.AuthGuardAdminOnly)

	// User routes
	ledgerRoute.Get("/", ledger.GetLedgers)
	ledgerRoute.Get("/:id", ledger.GetLedger)
	ledgerRoute.Post("/", ledger.CreateLedger)
	ledgerRoute.Put("/:id", ledger.UpdateLedger)
	ledgerRoute.Delete("/:id", ledger.DeleteLedger)
	ledgerRoute.Get("/account/:accountId", ledger.GetLedgersByAccount)
	ledgerRoute.Get("/date-range", ledger.GetLedgersByDateRange)

	// Admin routes
	adminLedgerRoute.Get("/", ledger.AdminGetAllLedgers)
	adminLedgerRoute.Get("/:id", ledger.AdminGetLedger)
	adminLedgerRoute.Get("/user/:userId", ledger.AdminGetLedgersByUser)
	adminLedgerRoute.Get("/account/:accountId", ledger.AdminGetLedgersByAccount)
	adminLedgerRoute.Get("/date-range", ledger.AdminGetLedgersByDateRange)

	log.Info("Ledger routes initialized")
}
