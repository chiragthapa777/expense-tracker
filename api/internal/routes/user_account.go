package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	useraccount "github.com/chiragthapa777/expense-tracker-api/internal/modules/user_account"
	"github.com/gofiber/fiber/v2"
)

func SetupUserAccountRoutes(app fiber.Router) {
	log := logger.GetLogger()
	userAccountRoute := app.Group("/user-account", middleware.AuthGuard)
	// adminBankRoute := app.Group("/admin/user-account", middleware.AuthGuardAdminOnly)

	userAccountModule := useraccount.NewUserAccountUserService()

	// user route
	userAccountRoute.Post("/", userAccountModule.CreateUserAccount)
	userAccountRoute.Get("/", userAccountModule.GetUserAccounts)
	userAccountRoute.Get("/:id", userAccountModule.GetUserAccountById)
	userAccountRoute.Put("/:id", userAccountModule.UpdateUserAccountById)
	userAccountRoute.Delete("/:id", userAccountModule.DeleteUserAccountById)

	// user admin routes
	// adminBankRoute.Post("/", bank.AdminCreateBank)
	// adminBankRoute.Get("/", bank.AdminGetBanks)
	// adminBankRoute.Get("/:id", bank.AdminGetBankById)
	// adminBankRoute.Put("/:id", bank.AdminUpdateBankById)
	// adminBankRoute.Delete("/:id", bank.AdminUpdateBankById)

	log.Info("Bank routes initialized")
}
