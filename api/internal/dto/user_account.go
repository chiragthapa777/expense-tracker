package dto

type UserAccountDto struct {
	AccountNumber *string `json:"accountNumber" validate:"omitempty,required_without=nil,min=3" trim:"true"`
	PhoneNumber   *string `json:"phoneNumber" validate:"omitempty,required_without=nil,min=3" trim:"true"`
	Name          string  `json:"name" validate:"required,min=3" trim:"true"`
	BankId        string  `json:"bankId" validate:"required,uuid"`
	IsActive      bool    `json:"isActive" validate:"required,boolean"`
}

type UserAccountCreateDto struct {
	Balance float64 `json:"balance" validate:"gte=0"`
	UserAccountDto
}

type UserAccountUpDto struct {
	UserAccountDto
}
