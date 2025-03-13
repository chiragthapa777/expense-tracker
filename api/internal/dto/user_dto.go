package dto

type UserProfileUpdateDto struct {
	FirstName string  `json:"firstName" validate:"required,min=3"`
	LastName  *string `json:"lastName" validate:"omitempty,required_without=nil,min=3"`
}

type UserPasswordUpdate struct {
	OldPassword string `json:"oldPassword" validate:"required,gte=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
	NewPassword string `json:"newPassword" validate:"required,gte=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}
