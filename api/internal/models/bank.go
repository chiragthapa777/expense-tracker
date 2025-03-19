package models

type Bank struct {
	BaseModel
	Name string `gorm:"type:varchar;not null" json:"name"`

	// associations
	UserAccounts *[]UserAccount `gorm:"foreignKey:bank_id;references:id" json:"userAccounts,omitempty"`
}

// TableName overrides the table name used by Bank to `banks`
func (Bank) TableName() string {
	return "banks"
}
