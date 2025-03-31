package dto

type BankCreateDto struct {
	Name string `json:"name" validate:"required,min=2" trim:"true"`
}
