package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"gorm.io/gorm"
)

type Option struct {
	Tx            *gorm.DB // for transaction
	PaginationDto *dto.PaginationQueryDto
	QueryBuilder  *gorm.DB // for custom additional filters and joins
	WithUserId    *string
	WithUserIdKey *string // defaults to user_id
}

type PaginationResult[T any] struct {
	MetaData types.ResponsePaginationMeta
	Data     []T
}
