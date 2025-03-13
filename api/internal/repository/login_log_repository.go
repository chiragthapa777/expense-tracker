package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
)

type LoginLogRepository struct {
	*BaseRepository[models.LoginLog]
}

func NewLoginLogRepository() *LoginLogRepository {
	return &LoginLogRepository{
		BaseRepository: NewBaseRepository[models.LoginLog](),
	}
}

// GetUserValidSortColumn returns valid sort columns for User.
func GetLoginLogValidSortColumn() map[string]string {
	return map[string]string{
		"id":        "id",
		"createdAt": "createdAt",
		"firstName": "first_name",
		"lastName":  "last_name",
	}
}

// FindWithPagination overrides the base method to specify User-specific search fields.
func (r *LoginLogRepository) FindWithPagination(option Option) (*PaginationResult[models.LoginLog], error) {
	searchFields := []string{}
	return r.BaseRepository.FindWithPagination(option, searchFields, GetLoginLogValidSortColumn())
}
