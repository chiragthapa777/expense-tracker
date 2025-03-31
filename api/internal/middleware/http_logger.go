package middleware

import (
	"fmt"

	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func HttpLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {

		logger := logger.GetLogger()

		logger.Info(fmt.Sprintf("[HTTP] method = %s, path = %s", c.Method(), c.Path()))
		return c.Next()
	}
}
