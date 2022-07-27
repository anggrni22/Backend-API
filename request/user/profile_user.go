package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegProfileUser struct {
	IdUser     string `json:"id_user" form:"id_user"`
	Token      string `json:"token" form:"token"`
	Identifier string `json:"identifier" form:"identifier"`
} 

func (p RegProfileUser) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.IdUser, validation.Required),
		validation.Field(&p.Identifier, validation.Required),
		validation.Field(&p.Token, validation.Required),
	)
}

type RegEditProfileUser struct {
	Email      string `json:"email" form:"email"`
	Phone      string `json:"phone" form:"phone"`
	FirstName  string `json:"first_name" form:"first_name"`
	LastName   string `json:"last_name" form:"last_name"`
	Specialist string `json:"specialist" form:"specialist"`
}

func (p RegEditProfileUser) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Email, validation.Required),
		validation.Field(&p.Phone, validation.Required),
		validation.Field(&p.FirstName, validation.Required),
		validation.Field(&p.LastName, validation.Required),
		validation.Field(&p.Specialist, validation.Required),
	)
}
