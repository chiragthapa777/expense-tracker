package repository

import (
	"time"

	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"gorm.io/gorm"
)

type LedgerRepository struct {
	*BaseRepository[models.Ledger]
}

func NewLedgerRepository() *LedgerRepository {
	return &LedgerRepository{
		BaseRepository: NewBaseRepository[models.Ledger](),
	}
}

// FindWithAssociations retrieves a ledger entry with all its associations
func (r *LedgerRepository) FindWithAssociations(id string, option Option) (*models.Ledger, error) {
	db := r.getDB(option)
	var ledger models.Ledger
	err := db.Preload("Account").Preload("User").
		Where("id = ?", id).First(&ledger).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &ledger, nil
}

// FindByAccountID retrieves all ledger entries for a specific account
func (r *LedgerRepository) FindByAccountID(accountID string, option Option) ([]models.Ledger, error) {
	db := r.getDB(option)
	var ledgers []models.Ledger
	err := db.Where("account_id = ?", accountID).Find(&ledgers).Error
	if err != nil {
		return nil, err
	}
	return ledgers, nil
}

// FindByDateRange retrieves ledger entries within a date range
func (r *LedgerRepository) FindByDateRange(startDate, endDate time.Time, option Option) ([]models.Ledger, error) {
	db := r.getDB(option)
	var ledgers []models.Ledger
	err := db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&ledgers).Error
	if err != nil {
		return nil, err
	}
	return ledgers, nil
}

// FindByUserID retrieves all ledger entries for a specific user
func (r *LedgerRepository) FindByUserID(userID string, option Option) ([]models.Ledger, error) {
	db := r.getDB(option)
	var ledgers []models.Ledger
	err := db.Where("user_id = ?", userID).Find(&ledgers).Error
	if err != nil {
		return nil, err
	}
	return ledgers, nil
}

// GetAccountBalance calculates the current balance for an account
func (r *LedgerRepository) GetAccountBalance(accountID string, option Option) (float64, error) {
	db := r.getDB(option)
	var balance float64
	err := db.Model(&models.Ledger{}).
		Where("account_id = ?", accountID).
		Select("COALESCE(SUM(credit - debit), 0)").
		Scan(&balance).Error
	if err != nil {
		return 0, err
	}
	return balance, nil
}
