package user

import (
	"errors"

	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AdminGetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if err := uuid.Validate(userId); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	userRepository := repository.NewUserRepository()

	user, err := userRepository.FindByID(userId, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if user == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  errors.New("user not found"),
			Status: fiber.StatusNotFound,
		})
	}

	return response.Send(c, types.ResponseOption{Data: user})
}

func AdminGetUsers(c *fiber.Ctx) error {
	query := dto.PaginationQueryDto{}
	if err := request.LoadAndValidateQuery(&query, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	userRepository := repository.NewUserRepository()

	queryBuilder := userRepository.GetQueryBuilder(repository.Option{})

	result, err := userRepository.FindWithPagination(repository.Option{PaginationDto: &query, QueryBuilder: queryBuilder})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{
		Data: result.Data,
		MetaData: &types.ResponseMetaData{
			PaginationMetaData: &result.MetaData,
		},
	})
}

func AdminBlockUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if err := uuid.Validate(userId); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	userRepository := repository.NewUserRepository()

	user, err := userRepository.FindByID(userId, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if user == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  errors.New("user not found"),
			Status: fiber.StatusNotFound,
		})
	}

	if user.BlockedAt != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user already blocked"), Status: fiber.StatusBadRequest})
	}

	if user.Role == models.UserRoleAdmin {
		return response.SendError(c, types.ErrorResponseOption{Error: errors.New("cannot block admin"), Status: fiber.StatusBadRequest})
	}

	if err = userRepository.BlockUser(userId, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: "user blocked"})
}

func AdminUnBlockUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if err := uuid.Validate(userId); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	userRepository := repository.NewUserRepository()

	user, err := userRepository.FindByID(userId, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if user == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  errors.New("user not found"),
			Status: fiber.StatusNotFound,
		})
	}

	if user.BlockedAt == nil {
		return response.SendError(c, types.ErrorResponseOption{Error: errors.New("user is not blocked"), Status: fiber.StatusBadRequest})
	}

	if err = userRepository.UnBlockUser(userId, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: "user unblocked"})
}

func AdminUpdateUser(c *fiber.Ctx) error {
	body := new(dto.UserProfileUpdateDto)
	if err := request.LoadAndValidateBody(body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	userId := c.Params("id")
	if err := uuid.Validate(userId); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	userRepository := repository.NewUserRepository()

	user, err := userRepository.FindByID(userId, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}
	if user == nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  errors.New("user not found"),
			Status: fiber.StatusNotFound,
		})
	}

	user.FirstName = body.FirstName
	user.LastName = body.LastName

	if err := userRepository.Update(user, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	return response.Send(c, types.ResponseOption{Data: user})
}
