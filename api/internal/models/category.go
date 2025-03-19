package models

type Category struct {
	BaseModel
	Name        string `gorm:"type:varchar;not null;unique" json:"name"`
	Icon        string `gorm:"type:varchar" json:"icon"`
	Color       string `gorm:"type:varchar;not null" json:"color"`
	Description string `gorm:"type:text;not null" json:"description"`
}

// TableName overrides the table name used by Category to `categories`
func (Category) TableName() string {
	return "categories"
}
