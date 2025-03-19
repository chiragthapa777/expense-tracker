package models

import "time"

type Ledger struct {
	BaseModel
	Debit         float64   `gorm:"type:decimal(10,2);not null;default:0" json:"debit"`
	Credit        float64   `gorm:"type:decimal(10,2);not null;default:0" json:"credit"`
	AccountID     *string   `gorm:"column:account_id;type:uuid" json:"accountId,omitempty"`
	TransactionID *string   `gorm:"column:transaction_id;type:varchar" json:"transactionId,omitempty"`
	UserID        string    `gorm:"column:user_id;type:uuid;not null" json:"userId"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	Date          time.Time `gorm:"type:date;not null" json:"date"`

	// associations
	Account *UserAccount `gorm:"foreignKey:account_id;references:id" json:"account,omitempty"`
	User    *User        `gorm:"foreignKey:user_id;references:id" json:"user,omitempty"`
}

// TableName overrides the table name used by Ledger to `ledgers`
func (Ledger) TableName() string {
	return "ledgers"
}
