package data

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

//Validate the fields of user model
func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 20)),
	)
}

//Validate the fields of context model
func (context Context) Validate() error {
	return validation.ValidateStruct(&context,
		validation.Field(&context.Txt, validation.Required),
	)
}
