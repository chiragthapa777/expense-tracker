package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app fiber.Router) {
	log := logger.GetLogger()
	userRoute := app.Group("/user", middleware.AuthGuard)
	adminUserRoute := app.Group("/admin/user", middleware.AuthGuardAdminOnly)

	// user route
	userRoute.Put("/update-profile", user.UpdateProfile)
	userRoute.Patch("/update-password", user.UpdatePassword)

	// user admin routes
	adminUserRoute.Get("/", user.AdminGetUsers)
	adminUserRoute.Get("/:id", user.AdminGetUser)
	adminUserRoute.Patch("/:id/block", user.AdminBlockUser)
	adminUserRoute.Patch("/:id/unblock", user.AdminUnBlockUser)
	adminUserRoute.Put("/:id/update", user.AdminUpdateUser)

	log.Info("User routes initialized")
}
