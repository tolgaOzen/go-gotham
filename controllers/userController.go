package controllers

import (
	"github.com/labstack/echo/v4"
	"gotham/helpers"
	"gotham/models/scopes"
	"gotham/requests"
	"gotham/services"
	"gotham/viewModels"
	"net/http"
)

type UserController struct {
	UserService services.IUserService
}

/**
* index
*
* @param echo.Context
* @return error
 */
func (u UserController) Index(c echo.Context) (err error) {

	request := new(scopes.Pagination)

	if err = c.Bind(request); err != nil {
		return
	}

	users, err := u.UserService.GetUsers(request, "name")
	if err != nil {
		return echo.ErrInternalServerError
	}

	count, err := u.UserService.GetUsersCount()
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(viewModels.Paginator{
		TotalRecord: int(count),
		Records:     users,
		Limit:       request.Limit,
		Page:        request.Page,
	}))
}

/**
* Show
*
* @param echo.Context
* @return error
 */
func (u UserController) Show(c echo.Context) (err error) {

	request := new(requests.UserShowRequest)

	if err = c.Bind(request); err != nil {
		return
	}

	v := request.Validate()

	if v != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"errors": v,
		})
	}


	user, err := u.UserService.GetUserByID(request.User)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(user))
}
