package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gotham/models/scopes"
)

type UserIndexRequest struct {
	validation.Validatable `json:"-" form:"-" query:"-"`

	/**
	 * PathParams
	 */
	PathParams struct{
	}

	/**
	 * QueryParams
	 */
	QueryParams struct{
		scopes.Pagination
	}

	/**
	 * Body
	 */
	Body struct{
	}
}

func (r UserIndexRequest) Validate() error {
	return nil
}

