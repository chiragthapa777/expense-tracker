package useraccount

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/auth"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type UserAccountUserController struct {
	service UserAccountService
}

func NewUserAccountUserController() *UserAccountUserController {
	u := &UserAccountUserController{
		service: *NewUserAccountService(),
	}
	return u
}

func (u UserAccountUserController) CreateUserAccount(c *fiber.Ctx) error {
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

	userAccount, err := u.service.CreateUserAccount(&body, *currentUser)

	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: userAccount})
}

// func (u UserAccountUserController) GetUserAccounts(c *fiber.Ctx) error {
// 	query := dto.PaginationQueryDto{}
// 	if err := request.LoadAndValidateQuery(&query, c); err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{Error: err})
// 	}

// 	userAccounts, err := u.userAccountRepository.FindWithPagination(repository.Option{PaginationDto: &query})

// 	if err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{
// 			Error: err,
// 		})
// 	}

// 	return response.Send(c, types.ResponseOption{Data: userAccounts.Data, MetaData: &types.ResponseMetaData{PaginationMetaData: &userAccounts.MetaData}})
// }

// func (u UserAccountUserController) GetUserAccountById(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if err := uuid.Validate(id); err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
// 	}
// 	userAccount, err := u.userAccountRepository.FindByID(id, repository.Option{})
// 	if err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{
// 			Error: err,
// 		})
// 	}

// 	return response.Send(c, types.ResponseOption{Data: userAccount})
// }

// func (u UserAccountUserController) UpdateUserAccountById(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if err := uuid.Validate(id); err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
// 	}

// 	body := dto.UserAccountCreateDto{}
// 	if err := request.LoadAndValidateBody(&body, c); err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{
// 			Status: fiber.StatusBadRequest,
// 			Error:  err,
// 		})
// 	}

// 	userAccount, err := u.userAccountRepository.FindByID(id, repository.Option{})
// 	if err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{
// 			Error: err,
// 		})
// 	}

// 	userAccount.Name = body.Name

// 	err = u.userAccountRepository.Update(userAccount, repository.Option{})
// 	if err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{
// 			Error: err,
// 		})
// 	}

// 	return response.Send(c, types.ResponseOption{Data: userAccount})
// }

// func (u UserAccountUserController) DeleteUserAccountById(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if err := uuid.Validate(id); err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
// 	}

// 	userAccount, err := u.userAccountRepository.FindByID(id, repository.Option{})
// 	if err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{
// 			Error: err,
// 		})
// 	}

// 	err = u.userAccountRepository.Delete(userAccount.ID, repository.Option{})
// 	if err != nil {
// 		return response.SendError(c, types.ErrorResponseOption{
// 			Error: err,
// 		})
// 	}
// 	return response.Send(c, types.ResponseOption{Data: userAccount})
// }
