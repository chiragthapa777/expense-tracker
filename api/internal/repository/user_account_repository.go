package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
)

type UserAccountRepository struct {
	*BaseRepository[models.UserAccount]
}

func NewUserAccountRepository() *UserAccountRepository {
	return &UserAccountRepository{
		BaseRepository: NewBaseRepository[models.UserAccount](),
	}
}

func GetUserAccountValidSortColumn() map[string]string {
	return map[string]string{
		"id":            "id",
		"name":          "name",
		"accountNumber": "account_number",
		"phoneNumber":   "phone_number",
		"lastName":      "last_name",
		"createdAt":     "created_at",
	}
}

func (r *UserAccountRepository) FindWithPagination(option Option) (*PaginationResult[models.UserAccount], error) {
	searchFields := []string{"name", "account_number", "phone_number"}
	return r.BaseRepository.FindWithPagination(option, searchFields, GetUserAccountValidSortColumn())
}

func (r *UserAccountRepository) FindByName(name string, option Option) ([]models.UserAccount, error) {
	db := r.getDB(option)
	var banks []models.UserAccount
	err := db.Where("name ILIKE ?", name).Find(&banks).Error
	return banks, err
}
