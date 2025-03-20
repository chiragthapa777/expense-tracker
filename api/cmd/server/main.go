package main

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/config"
	"github.com/chiragthapa777/expense-tracker-api/internal/database"
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/middleware"
	"github.com/chiragthapa777/expense-tracker-api/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log := logger.GetLogger()
	log.Info("test1")

	config := config.GetConfig()

	database.InitDB()

	app := fiber.New()

	app.Use(middleware.CORSMiddleware())

	routes.SetUpRoutes(app)

	log.Info("Server starting on port " + config.Port)
	err := app.Listen(":" + config.Port)
	if err != nil {
		log.Errorf("Failed to start server: %v", err)
		panic("Server failed to start")
	}
}
