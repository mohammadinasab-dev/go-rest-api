package data

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

//Validate is
func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 20)),
	)
}

//Validate is
func (context Context) Validate() error {
	return validation.ValidateStruct(&context,
		validation.Field(&context.Txt, validation.Required),
	)
}
