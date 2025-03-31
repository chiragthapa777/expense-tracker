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

	// user route
	bankRoute.Get("/", bank.GetBanks)
	bankRoute.Get("/:id", bank.GetBankById)

	// user admin routes
	adminBankRoute.Post("/", bank.AdminCreateBank)
	adminBankRoute.Get("/", bank.AdminGetBanks)
	adminBankRoute.Get("/:id", bank.AdminGetBankById)
	adminBankRoute.Put("/:id", bank.AdminUpdateBankById)
	adminBankRoute.Delete("/:id", bank.AdminUpdateBankById)

	log.Info("Bank routes initialized")
}
