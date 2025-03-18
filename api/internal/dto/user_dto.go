package dto

type UserProfileUpdateDto struct {
	FirstName     string  `json:"firstName" validate:"required,min=3"`
	LastName      *string `json:"lastName" validate:"omitempty,required_without=nil,min=3"`
	ProfileId     *string `json:"profileId" validate:"uuid"`
	RemoveProfile bool    `json:"removeProfile" validate:"boolean"`
}

type UserPasswordUpdate struct {
	OldPassword string `json:"oldPassword" validate:"required,gte=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
	NewPassword string `json:"newPassword" validate:"required,gte=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}
