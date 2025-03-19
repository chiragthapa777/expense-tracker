package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/user_account"
	"github.com/gofiber/fiber/v2"
)

func SetupUserAccountRoutes(app fiber.Router) {
	log := logger.GetLogger()
	accountRoute := app.Group("/account", middleware.AuthGuard)
	adminAccountRoute := app.Group("/admin/account", middleware.AuthGuardAdminOnly)

	// User routes
	accountRoute.Get("/", user_account.GetUserAccounts)
	accountRoute.Get("/:id", user_account.GetUserAccount)
	accountRoute.Post("/", user_account.CreateUserAccount)
	accountRoute.Put("/:id", user_account.UpdateUserAccount)
	accountRoute.Delete("/:id", user_account.DeleteUserAccount)
	accountRoute.Get("/:id/balance", user_account.GetAccountBalance)
	accountRoute.Get("/:id/ledgers", user_account.GetAccountLedgers)

	// Admin routes
	adminAccountRoute.Get("/", user_account.GetAllAccounts)
	adminAccountRoute.Get("/:id", user_account.GetAccount)
	adminAccountRoute.Post("/", user_account.CreateAccount)
	adminAccountRoute.Put("/:id", user_account.UpdateAccount)
	adminAccountRoute.Delete("/:id", user_account.DeleteAccount)
	adminAccountRoute.Get("/users/:userId/accounts", user_account.GetUserAccountsAdmin)

	log.Info("User account routes initialized")
}
