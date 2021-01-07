package controllers

import (
	"github.com/labstack/echo/v4"
	"gotham/app/container/dic"
	"gotham/helpers"
	"gotham/models"
	"gotham/models/accessories"
	"gotham/models/scopes"
	"gotham/requests"
	"net/http"
)

type UserController struct{}

/**
* index
*
* @param echo.Context
* @return error
 */
func (UserController) Index(c echo.Context) (err error) {

	request := new(requests.Pagination)

	if err = c.Bind(request); err != nil {
		return
	}

	var users []models.User

	if err := dic.Db(c.Request()).Scopes(scopes.Paginate(request, models.User{}, "name")).Find(&users).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(accessories.Paginator{
		TotalRecord: 0,
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
func (UserController) Show(c echo.Context) (err error) {

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

	var user models.User

	if err := dic.Db(c.Request()).Where("verified = ?", request.Verified).First(&user, request.User).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(user))
}
