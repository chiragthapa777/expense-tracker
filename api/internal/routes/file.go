package routes

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/file"
	"github.com/gofiber/fiber/v2"
)

func SetupFileRoutes(app fiber.Router) {
	log := logger.GetLogger()
	fileRouterGroup := app.Group("/file", middleware.AuthGuard)

	fileRouterGroup.Post("/upload-images", middleware.FileCheck(middleware.ImageFileCheckConfig), file.UploadImages)
	fileRouterGroup.Delete("/:ids", file.DeleteFiles)

	log.Info("File routes initialized")
}
