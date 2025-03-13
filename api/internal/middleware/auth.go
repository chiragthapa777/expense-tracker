package middleware

import (
	"errors"
	"slices"
	"strings"

	"github.com/chiragthapa777/expense-tracker-api/internal/constant"
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/chiragthapa777/expense-tracker-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddlewareOption struct {
	Roles               *[]models.UserRoleEnum
	AllowPasswordNotSet bool
}

func AuthMiddleware(options AuthMiddlewareOption) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log := logger.GetLogger()
		userRepository := repository.NewUserRepository()
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			log.Warn("Missing or invalid Authorization header")
			return response.SendError(c, types.ErrorResponseOption{
				Status: fiber.StatusUnauthorized,
				Error:  errors.New("unauthorized"),
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			log.Warnf("Invalid JWT: %v", err)
			return response.SendError(c, types.ErrorResponseOption{
				Status: fiber.StatusUnauthorized,
				Error:  errors.New("unauthorized"),
			})
		}

		user, err := userRepository.FindByID(claims.UserID, repository.Option{})

		if err != nil {
			return response.SendError(c, types.ErrorResponseOption{
				Status: fiber.StatusUnauthorized,
				Error:  errors.New("cannot find user"),
			})
		}

		if user.BlockedAt != nil {
			return response.SendError(c, types.ErrorResponseOption{
				Status: fiber.StatusUnauthorized,
				Error:  errors.New("user blocked, contact to admin"),
				Code:   constant.ResponseCodeUserBlocked,
			})
		}

		if !user.IsPasswordSetByUser && !options.AllowPasswordNotSet {
			return response.SendError(c, types.ErrorResponseOption{
				Status: fiber.StatusBadRequest,
				Error:  errors.New("unauthorized"),
				Code:   constant.ResponseCodePasswordNotSet,
			})
		}

		if options.Roles != nil && !slices.Contains(*options.Roles, user.Role) {
			return response.SendError(c, types.ErrorResponseOption{
				Status: fiber.StatusUnauthorized,
				Error:  errors.New("not permitted"),
			})
		}

		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}

var AuthGuard = AuthMiddleware(AuthMiddlewareOption{})
var AuthGuardWithAllowPasswordNotSet = AuthMiddleware(AuthMiddlewareOption{AllowPasswordNotSet: true})
var AuthGuardAdminOnly = AuthMiddleware(AuthMiddlewareOption{Roles: &[]models.UserRoleEnum{models.UserRoleAdmin}})
