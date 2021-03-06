package requests

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginRequest struct {
	validation.Validatable `json:"-" form:"-" query:"-"`

	/**
    * BODY
    */
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

func (r LoginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, validation.Length(4, 50), is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 50)),
	)
}
