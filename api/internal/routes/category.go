package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/category"
	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(app fiber.Router) {
	log := logger.GetLogger()
	categoryRoute := app.Group("/category", middleware.AuthGuard)
	adminCategoryRoute := app.Group("/admin/category", middleware.AuthGuardAdminOnly)

	// User routes
	categoryRoute.Get("/", category.GetCategories)
	categoryRoute.Get("/:id", category.GetCategory)
	categoryRoute.Get("/color/:color", category.GetCategoriesByColor)

	// Admin routes
	adminCategoryRoute.Post("/", category.CreateCategory)
	adminCategoryRoute.Put("/:id", category.UpdateCategory)
	adminCategoryRoute.Delete("/:id", category.DeleteCategory)

	log.Info("Category routes initialized")
}
