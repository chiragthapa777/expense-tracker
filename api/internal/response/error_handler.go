package response

import (
	"errors"

	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

// This will catch all the error which has not been handled
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Return error response from out own handler function
	return SendError(ctx, types.ErrorResponseOption{
		Status: code,
		Error:  err,
	})
}
