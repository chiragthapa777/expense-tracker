package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/auth"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app fiber.Router) {
	log := logger.GetLogger()
	authGroup := app.Group("/auth")

	authGroup.Post("/register", auth.Register)
	authGroup.Post("/login", auth.Login)
	authGroup.Get("/me", middleware.AuthGuardWithAllowPasswordNotSet, auth.GetCurrentUser)
	authGroup.Post("/set-password", middleware.AuthGuardWithAllowPasswordNotSet, auth.SetPassword)

	log.Info("Auth routes initialized")
}
