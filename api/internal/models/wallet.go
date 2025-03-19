package models

type Wallet struct {
	BaseModel
	Name string `gorm:"type:varchar;not null" json:"name"`

	// associations
	UserAccounts *[]UserAccount `gorm:"foreignKey:wallet_id;references:id" json:"userAccounts,omitempty"`
}

// TableName overrides the table name used by Wallet to `wallets`
func (Wallet) TableName() string {
	return "wallets"
}
