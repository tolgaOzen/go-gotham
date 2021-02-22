package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserShowRequest struct {
	validation.Validatable `json:"-" form:"-" query:"-"`

	/**
	 * PATH
	 */
	User uint `path:"user" json:"-" form:"-" query:"-" xml:"-"`

	/**
	 * BODY
	 */
	Verified int `json:"verified" form:"verified" query:"verified"`
}

func (r UserShowRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Verified, validation.Required, validation.Min(0), validation.Max(1)),
	)
}
