package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gotham/utils"
)

type UserIndexRequest struct {
	validation.Validatable `json:"-" form:"-" query:"-"`

	/**
	 * PathParams
	 */
	PathParams struct{}

	/**
	 * QueryParams
	 */
	QueryParams struct {
		utils.Order
		utils.Pagination
	}

	/**
	 * Body
	 */
	Body struct{}
}

func (r UserIndexRequest) Validate() error {
	return nil
}
