package response

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func formatValidationError(err validator.FieldError) string {
	fieldName := err.Field() // Uses JSON tag due to RegisterTagNameFunc
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldName)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fieldName, err.Param())
	case "gte":
		return fmt.Sprintf("%s must be at least %s", fieldName, err.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldName)
	case "containsany":
		switch err.Param() {
		case "ABCDEFGHIJKLMNOPQRSTUVWXYZ":
			return fmt.Sprintf("%s must contain at least one uppercase letter", fieldName)
		case "abcdefghijklmnopqrstuvwxyz":
			return fmt.Sprintf("%s must contain at least one lowercase letter", fieldName)
		case "0123456789":
			return fmt.Sprintf("%s must contain at least one number", fieldName)
		default:
			return fmt.Sprintf("%s has invalid characters", fieldName)
		}
	default:
		return fmt.Sprintf("%s failed validation: %s", fieldName, err.Tag())
	}
}
