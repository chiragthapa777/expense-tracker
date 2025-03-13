package models

type MethodEnum string

const (
	MethodEnumPASSWORD MethodEnum = "PASSWORD"
	MethodEnumOAUTH    MethodEnum = "OAUTH"
)

type LoginLog struct {
	BaseModel
	UserID string     `gorm:"column:user_id;type:uuid;not null" json:"userId"`
	Method MethodEnum `gorm:"type:varchar(10);not null" json:"method"`

	// associations
	User *User `gorm:"foreignKey:user_id;references:id;constraint:OnDelete:CASCADE;" json:"user,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (LoginLog) TableName() string {
	return "login_logs"
}
