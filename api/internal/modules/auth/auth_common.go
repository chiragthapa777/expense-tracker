package auth

import (
	"errors"

	"github.com/chiragthapa777/expense-tracker-api/internal/constant"
	"github.com/chiragthapa777/expense-tracker-api/internal/database"
	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/chiragthapa777/expense-tracker-api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	userRepository := repository.NewUserRepository()

	body := new(dto.AuthRegisterDto)
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  err,
		})
	}

	emailUser, err := userRepository.FindByEmail(body.Email)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Status: fiber.StatusBadRequest, Error: err})
	}
	if emailUser != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  errors.New("email already exists"),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	user := &models.User{
		FirstName:           body.FirstName,
		LastName:            body.LastName,
		Email:               &body.Email,
		Password:            string(hashedPassword),
		IsPasswordSetByUser: true,
	}

	if err := userRepository.Create(user, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Status: fiber.StatusCreated,
		Data:   map[string]any{"user": user},
	})
}

func Login(c *fiber.Ctx) error {
	userRepository := repository.NewUserRepository()
	loginLogRepository := repository.NewLoginLogRepository()

	body := new(dto.AuthLoginDto)
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	foundUser, err := userRepository.FindByEmail(body.Email)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if foundUser == nil {
		return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user not found"), Status: fiber.StatusBadRequest})
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(body.Password))
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: errors.New("invalid credentials"), Status: fiber.StatusBadRequest})
	}

	if foundUser.BlockedAt != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  errors.New("user blocked, contact to admin"),
			Code:   constant.ResponseCodeUserBlocked,
		})
	}

	token, err := utils.GenerateJWT(foundUser.ID)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	tx := database.DB.Begin()

	if err = loginLogRepository.Create(&models.LoginLog{
		UserID: foundUser.ID,
		Method: models.MethodEnumPASSWORD,
	}, repository.Option{Tx: tx}); err != nil {
		tx.Rollback()
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	tx.Commit()

	return response.Send(c, types.ResponseOption{
		Status: fiber.StatusCreated,
		Data:   map[string]any{"token": token, "user": foundUser},
	})
}

func CurrentUser(c *fiber.Ctx) (*models.User, error) {
	userRepository := repository.NewUserRepository()
	userID := c.Locals("user_id").(string)

	if userID == "" {
		return nil, repository.ErrRecordNotFound
	}

	user, err := userRepository.FindByIdWithJoins(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, repository.ErrRecordNotFound
	}

	return user, nil
}

func GetCurrentUser(c *fiber.Ctx) error {
	user, err := CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	return response.Send(c, types.ResponseOption{Data: user})
}

func SetPassword(c *fiber.Ctx) error {
	user, err := CurrentUser(c)
	userRepo := repository.NewUserRepository()
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	if user.IsPasswordSetByUser {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  errors.New("password is already set"),
			Status: fiber.StatusBadRequest,
		})
	}

	body := new(dto.SetPasswordDto)
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  err,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	if err = userRepo.UpdatePassword(user.ID, string(hashedPassword), repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: user})
}
