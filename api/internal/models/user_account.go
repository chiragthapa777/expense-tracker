package models

type UserAccount struct {
	BaseModel
	BankID string `gorm:"column:bank_id;type:uuid" json:"bankId,omitempty"`
	// WalletID      *string `gorm:"column:wallet_id;type:uuid" json:"walletId,omitempty"`
	UserID        string  `gorm:"column:user_id;type:uuid;not null" json:"userId"`
	AccountNumber *string `gorm:"column:account_number;type:varchar" json:"accountNumber,omitempty"`
	PhoneNumber   *string `gorm:"column:phone_number;type:varchar" json:"phoneNumber,omitempty"`
	Name          string  `gorm:"type:varchar;not null" json:"name"`
	Balance       float64 `gorm:"type:decimal(10,2);not null;default:0" json:"balance"`
	IsActive      bool    `gorm:"not null;default:true" json:"isActive"`

	// associations
	Bank *Bank `gorm:"foreignKey:bank_id;references:id" json:"bank,omitempty"`
	// Wallet  *Wallet   `gorm:"foreignKey:wallet_id;references:id" json:"wallet,omitempty"`
	User *User `gorm:"foreignKey:user_id;references:id" json:"user,omitempty"`
	// Ledgers *[]Ledger `gorm:"foreignKey:account_id;references:id" json:"ledgers,omitempty"`
}

// TableName overrides the table name used by UserAccount to `user_accounts`
func (UserAccount) TableName() string {
	return "user_accounts"
}
