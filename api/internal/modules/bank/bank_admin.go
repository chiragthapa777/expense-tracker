package bank

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
	"gorm.io/gorm"
)

func AdminCreateBank(c *fiber.Ctx) error {
	body := dto.BankCreateDto{}
	if err := request.LoadAndValidateBody(&body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  err,
		})
	}

	bankRepo := repository.NewBankRepository()

	_, err := bankRepo.FindByName(
		body.Name,
		repository.Option{},
	)

	if err != gorm.ErrRecordNotFound {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  errors.New("bank already exists"),
		})
	}

	bank := &models.Bank{
		Name: body.Name,
	}

	if err := bankRepo.Create(bank, repository.Option{}); err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: bank})
}

func AdminGetBanks(c *fiber.Ctx) error {
	query := dto.PaginationQueryDto{}
	if err := request.LoadAndValidateQuery(&query, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err})
	}

	bankRepo := repository.NewBankRepository()

	banks, err := bankRepo.FindWithPagination(repository.Option{PaginationDto: &query})

	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: banks.Data, MetaData: &types.ResponseMetaData{PaginationMetaData: &banks.MetaData}})
}

func AdminGetBankById(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uuid.Validate(id); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}
	bankRepo := repository.NewBankRepository()
	bank, err := bankRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: bank})
}

func AdminUpdateBankById(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uuid.Validate(id); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	body := dto.BankCreateDto{}
	if err := request.LoadAndValidateBody(&body, c); err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Status: fiber.StatusBadRequest,
			Error:  err,
		})
	}

	bankRepo := repository.NewBankRepository()

	bank, err := bankRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	foundBanks, err := bankRepo.FindByNameExceptId(bank.Name, id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	if len(foundBanks) > 0 {
		return response.SendError(c, types.ErrorResponseOption{
			Error:  errors.New("bank already exists"),
			Status: fiber.StatusBadRequest,
		})
	}

	bank.Name = body.Name

	err = bankRepo.Update(bank, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	return response.Send(c, types.ResponseOption{Data: bank})
}

func AdminDeleteBankById(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uuid.Validate(id); err != nil {
		return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
	}

	bankRepo := repository.NewBankRepository()

	bank, err := bankRepo.FindByID(id, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}

	err = bankRepo.Delete(bank.ID, repository.Option{})
	if err != nil {
		return response.SendError(c, types.ErrorResponseOption{
			Error: err,
		})
	}
	return response.Send(c, types.ResponseOption{Data: bank})
}
