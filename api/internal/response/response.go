package response

import (
	"errors"
	"strings"

	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SendError(c *fiber.Ctx, option types.ErrorResponseOption) error {
	log := logger.GetLogger()
	resp := types.Response{Success: false}
	status := fiber.StatusInternalServerError // Default status

	// Set error and status based on input
	if option.Error != nil {
		if validationErrors, ok := option.Error.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, ve := range validationErrors {
				errorMessages = append(errorMessages, formatValidationError(ve))
			}
			resp.Error = "Validation Error: " + strings.Join(errorMessages, ", ")
			status = fiber.StatusBadRequest
		} else if errors.Is(option.Error, repository.ErrRecordNotFound) {
			resp.Error = option.Error.Error()
			status = fiber.StatusNotFound
		} else {
			resp.Error = option.Error.Error()
			status = fiber.StatusInternalServerError
		}
		log.Errorf("Response error: %v", option.Error)
	} else {
		resp.Error = "An unexpected error occurred"
		log.Errorf("No error provided in SendErrorResponse")
	}

	// Override status if explicitly provided
	if option.Status != 0 {
		status = option.Status
	}

	if option.Code != "" {
		resp.Code = option.Code
	}

	return c.Status(status).JSON(resp)
}

func Send(c *fiber.Ctx, options types.ResponseOption) error {
	log := logger.GetLogger()
	resp := types.Response{
		Success: true,
		Data:    options.Data,
	}
	status := fiber.StatusOK

	if options.Status != 0 {
		status = options.Status
	}

	if options.MetaData != nil {
		resp.MetaData = options.MetaData
	}

	log.Info("Response sent successfully")
	return c.Status(status).JSON(resp)
}
