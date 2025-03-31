package useraccount

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/auth"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserAccountUserService struct {
	userAccountRepository repository.UserAccountRepository
}

func NewUserAccountUserService() *UserAccountUserService {
	u := &UserAccountUserService{
		userAccountRepository: *repository.NewUserAccountRepository(),
	}
	return u
}

func (u UserAccountUserService) CreateUserAccount(c *fiber.Ctx) error {
	body := dto.UserAccountCreateDto{}
	if err := request.LoadAndValidateBody(&body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  err,
		})
	}

	currentUser, err := auth.CurrentUser(c)
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	userAccount := &models.UserAccount{
		Name:          body.Name,
		AccountNumber: body.AccountNumber,
		PhoneNumber:   body.AccountNumber,
		Balance:       body.Balance,
		IsActive:      body.IsActive,
		UserID:        currentUser.ID,
		BankID:        body.BankId,
	}

	if err := u.userAccountRepository.Create(userAccount, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: userAccount})
}

func (u UserAccountUserService) GetUserAccounts(c *fiber.Ctx) error {
	query := dto.PaginationQueryDto{}
	if err := request.LoadAndValidateQuery(&query, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	userAccounts, err := u.userAccountRepository.FindWithPagination(repository.Option{PaginationDto: &query})

	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: userAccounts.Data, MetaData: &types.ResponseMetaData{PaginationMetaData: &userAccounts.MetaData}})
}

func (u UserAccountUserService) GetUserAccountById(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uuid.Validate(id); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}
	userAccount, err := u.userAccountRepository.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: userAccount})
}

func (u UserAccountUserService) UpdateUserAccountById(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uuid.Validate(id); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	body := dto.UserAccountCreateDto{}
	if err := request.LoadAndValidateBody(&body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  err,
		})
	}

	userAccount, err := u.userAccountRepository.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	userAccount.Name = body.Name

	err = u.userAccountRepository.Update(userAccount, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: userAccount})
}

func (u UserAccountUserService) DeleteUserAccountById(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uuid.Validate(id); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	userAccount, err := u.userAccountRepository.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	err = u.userAccountRepository.Delete(userAccount.ID, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}
	return response.Send(c, types.ResponseOption{Data: userAccount})
}
