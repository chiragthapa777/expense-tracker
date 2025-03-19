package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"gorm.io/gorm"
)

type WalletRepository struct {
	*BaseRepository[models.Wallet]
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{
		BaseRepository: NewBaseRepository[models.Wallet](),
	}
}

// FindWithUserAccounts retrieves a wallet with its associated user accounts
func (r *WalletRepository) FindWithUserAccounts(id string, option Option) (*models.Wallet, error) {
	db := r.getDB(option)
	var wallet models.Wallet
	err := db.Preload("UserAccounts").Where("id = ?", id).First(&wallet).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &wallet, nil
}

// FindAllWithUserAccounts retrieves all wallets with their associated user accounts
func (r *WalletRepository) FindAllWithUserAccounts(option Option) ([]models.Wallet, error) {
	db := r.getDB(option)
	var wallets []models.Wallet
	err := db.Preload("UserAccounts").Find(&wallets).Error
	if err != nil {
		return nil, err
	}
	return wallets, nil
}
