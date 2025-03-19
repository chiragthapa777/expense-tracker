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
