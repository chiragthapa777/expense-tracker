package bank

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/request"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetBanks(c *fiber.Ctx) error {
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

func GetBankById(c *fiber.Ctx) error {
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
