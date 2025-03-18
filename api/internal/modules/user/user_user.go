package user

import (
	"errors"

	"github.com/chiragthapa777/expense-tracker-api/internal/database"
	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/auth"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UpdateProfile(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	body := new(dto.UserProfileUpdateDto)
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	userRepository := repository.NewUserRepository()

	currentUser.FirstName = body.FirstName
	currentUser.LastName = body.LastName

	tx := database.DB.Begin()

	if err := userRepository.Update(currentUser, repository.Option{Tx: tx}); err != nil {
		tx.Rollback()
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	if err := UpdateProfilePicture(*body, *currentUser, repository.Option{Tx: tx}); err != nil {
		tx.Rollback()
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	tx.Commit()

	return response.Send(c, types.ResponseOption{})
}

func UpdatePassword(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	body := new(dto.UserPasswordUpdate)
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(body.OldPassword))
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: errors.New("wrong password"), Status: fiber.StatusBadRequest})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	userRepository := repository.NewUserRepository()

	if err := userRepository.UpdatePassword(currentUser.ID, string(hashedPassword), repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{})
	}

	return response.Send(c, types.ResponseOption{Data: "password updated successfully"})
}
