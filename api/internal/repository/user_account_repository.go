package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"gorm.io/gorm"
)

type UserAccountRepository struct {
	*BaseRepository[models.UserAccount]
}

func NewUserAccountRepository() *UserAccountRepository {
	return &UserAccountRepository{
		BaseRepository: NewBaseRepository[models.UserAccount](),
	}
}

// FindWithAssociations retrieves a user account with all its associations
func (r *UserAccountRepository) FindWithAssociations(id string, option Option) (*models.UserAccount, error) {
	db := r.getDB(option)
	var account models.UserAccount
	err := db.Preload("Bank").Preload("Wallet").Preload("User").Preload("Ledgers").
		Where("id = ?", id).First(&account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

// FindByUserID retrieves all accounts for a specific user
func (r *UserAccountRepository) FindByUserID(userID string, option Option) ([]models.UserAccount, error) {
	db := r.getDB(option)
	var accounts []models.UserAccount
	err := db.Preload("Bank").Preload("Wallet").
		Where("user_id = ?", userID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// UpdateBalance updates the balance of a user account
func (r *UserAccountRepository) UpdateBalance(id string, amount float64, option Option) error {
	db := r.getDB(option)
	return db.Model(&models.UserAccount{}).
		Where("id = ?", id).
		Update("balance", gorm.Expr("balance + ?", amount)).Error
}
