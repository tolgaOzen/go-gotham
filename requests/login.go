package requests

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginRequest struct {
	validation.Validatable `json:"-" form:"-" query:"-"`

	/**
	 * PathParams
	 */
	PathParams struct{}

	/**
	 * QueryParams
	 */
	QueryParams struct{}

	/**
	 * Body
	 */
	Body struct {
		Email    string `json:"email" form:"email" xml:"email"`
		Password string `json:"password" form:"password" xml:"password"`
	}
}

func (r LoginRequest) Validate() error {
	return validation.ValidateStruct(&r.Body,
		validation.Field(&r.Body.Email, validation.Required, validation.Length(4, 50), is.Email),
		validation.Field(&r.Body.Password, validation.Required, validation.Length(8, 50)),
	)
}
