package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	log := logger.GetLogger()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		log.Info("Received request to /")
		return c.SendString("Hello, World!")
	})

	// SetupAuthRoutes(v1)
	// SetupUserRoutes(v1)
	// SetupFileRoutes(v1)

	// // New routes
	// SetupBankRoutes(v1)
	// SetupWalletRoutes(v1)
	// SetupUserAccountRoutes(v1)
	// SetupCategoryRoutes(v1)
	// SetupLedgerRoutes(v1)
}
