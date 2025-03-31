package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"gorm.io/gorm"
)

type BankRepository struct {
	*BaseRepository[models.Bank]
}

func NewBankRepository() *BankRepository {
	return &BankRepository{
		BaseRepository: NewBaseRepository[models.Bank](),
	}
}

func GetBankValidSortColumn() map[string]string {
	return map[string]string{
		"id":        "id",
		"name":      "name",
		"lastName":  "last_name",
		"createdAt": "created_at",
	}
}
func (r *BankRepository) FindWithPagination(option Option) (*PaginationResult[models.Bank], error) {
	searchFields := []string{"name"}
	return r.BaseRepository.FindWithPagination(option, searchFields, GetBankValidSortColumn())
}

// FindWithUserAccounts retrieves a bank with its associated user accounts
func (r *BankRepository) FindWithUserAccounts(id string, option Option) (*models.Bank, error) {
	db := r.getDB(option)
	var bank models.Bank
	err := db.Preload("UserAccounts").Where("id = ?", id).First(&bank).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &bank, nil
}

// FindAllWithUserAccounts retrieves all banks with their associated user accounts
func (r *BankRepository) FindAllWithUserAccounts(option Option) ([]models.Bank, error) {
	db := r.getDB(option)
	var banks []models.Bank
	err := db.Preload("UserAccounts").Find(&banks).Error
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func (r *BankRepository) FindByName(name string, option Option) (*models.Bank, error) {
	db := r.getDB(option)
	var bank models.Bank
	err := db.Where("name ILIKE ?", name).First(&bank).Error
	return &bank, err
}

func (r *BankRepository) FindByNameExceptId(name string, id string, option Option) ([]models.Bank, error) {
	db := r.getDB(option)
	var banks []models.Bank = make([]models.Bank, 0)
	err := db.Where("name ILIKE ? and id <> ?", name, id).Find(&banks).Error
	return banks, err
}
