package useraccount

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserAccountService struct {
	userAccountRepository repository.UserAccountRepository
	bankRepository        repository.BankRepository
}

func NewUserAccountService() *UserAccountService {
	u := &UserAccountService{
		userAccountRepository: *repository.NewUserAccountRepository(),
		bankRepository:        *repository.NewBankRepository(),
	}
	return u
}

func (u UserAccountService) CreateUserAccount(body *dto.UserAccountCreateDto, user models.User) (*models.UserAccount, error) {

	_, err := u.bankRepository.FindByID(body.BankId, repository.Option{})
	if err != nil {
		return nil, err
	}

	userAccount := &models.UserAccount{
		Name:          body.Name,
		AccountNumber: body.AccountNumber,
		PhoneNumber:   body.AccountNumber,
		Balance:       body.Balance,
		IsActive:      body.IsActive,
		UserID:        user.ID,
		BankID:        body.BankId,
	}

	err = u.userAccountRepository.Create(userAccount, repository.Option{})

	return userAccount, err
}

func (u UserAccountService) GetUserAccounts(c *fiber.Ctx) error {
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

func (u UserAccountService) GetUserAccountById(c *fiber.Ctx) error {
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

func (u UserAccountService) UpdateUserAccountById(c *fiber.Ctx) error {
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

func (u UserAccountService) DeleteUserAccountById(c *fiber.Ctx) error {
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
