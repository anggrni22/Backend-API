package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegRegister struct {
	FullName  string `json:"full_name" form:"full_name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Identifier int    `json:"id_identifier" form:"id_identifier"`
}

func (reg RegRegister) Validate() error {
	return validation.ValidateStruct(&reg,
		validation.Field(&reg.FullName, validation.Required, validation.Length(3, 100)),
		validation.Field(&reg.Email, validation.Required, is.Email),
		validation.Field(&reg.Password, validation.Required, validation.Length(6, 50)),
	)
}
