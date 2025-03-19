package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/wallet"
	"github.com/gofiber/fiber/v2"
)

func SetupWalletRoutes(app fiber.Router) {
	log := logger.GetLogger()
	walletRoute := app.Group("/wallet", middleware.AuthGuard)
	adminWalletRoute := app.Group("/admin/wallet", middleware.AuthGuardAdminOnly)

	// user routes
	walletRoute.Get("/", wallet.GetWallets)
	walletRoute.Get("/:id", wallet.GetWallet)
	walletRoute.Get("/:id/accounts", wallet.GetWalletAccounts)

	// admin routes
	adminWalletRoute.Post("/", wallet.CreateWallet)
	adminWalletRoute.Put("/:id", wallet.UpdateWallet)
	adminWalletRoute.Delete("/:id", wallet.DeleteWallet)

	log.Info("Wallet routes initialized")
}
