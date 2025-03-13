package dto

type AuthRegisterDto struct {
	FirstName string  `json:"firstName" validate:"required,min=3"`
	LastName  *string `json:"lastName" validate:"omitempty,required_without=nil,min=3"`
	AuthLoginDto
}

type AuthLoginDto struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,gte=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}

type SetPasswordDto struct {
	Password string `json:"password" validate:"required,gte=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}
