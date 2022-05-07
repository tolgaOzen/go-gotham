package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserShowRequest struct {
	validation.Validatable `json:"-" form:"-" query:"-"`

	/**
	 * PathParams
	 */
	PathParams struct {
		User uint `param:"user"`
	}

	/**
	 * QueryParams
	 */
	QueryParams struct{}

	/**
	 * Body
	 */
	Body struct{}
}

func (r UserShowRequest) Validate() error {
	return nil
}
