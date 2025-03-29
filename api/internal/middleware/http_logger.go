package middleware

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func HttpLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := logger.GetLogger()
		logger.Info("Testerrrrrrrrrrr")
		return c.Next()
	}
}
