package models

import (
	"time"
)

type UserRoleEnum string

const (
	UserRoleUser  UserRoleEnum = "USER"
	UserRoleAdmin UserRoleEnum = "ADMIN"
)

type User struct {
	BaseModel
	FirstName           string     `gorm:"column:first_name;type:varchar;not null" json:"firstName"`
	LastName            *string    `gorm:"column:last_name;type:varchar" json:"lastName"`
	Email               *string    `gorm:"type:varchar" json:"email"`
	EmailVerifiedAt     *time.Time `gorm:"column:email_verified_at" json:"emailVerifiedAt"`
	BlockedAt           *time.Time `gorm:"column:blocked_at" json:"blockedAt"`
	Password            string     `gorm:"type:varchar;not null" json:"-"`
	IsPasswordSetByUser bool       `gorm:"type:boolean;not null;column:is_password_set_by_user" json:"isFirstLoggedIn"`
	Role                UserRoleEnum

	// associations
	LoginLogs *[]LoginLog `gorm:"foreignKey:user_id;references:id" json:"loginLogs,omitempty"`
	Profile   *File       `gorm:"foreignKey:UserProfileID;references:id" json:"profile"`
}
