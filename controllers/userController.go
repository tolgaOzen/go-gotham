package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gotham/config"
	"gotham/models"
	"gotham/models/scopes"
	"gotham/policies"
	"gotham/requests"
	"gotham/services"
	"gotham/viewModels"
	"net/http"
)

type UserController struct {
	UserService services.IUserService

	UserPolicy policies.IUserPolicy
}

/**
* index
*
* @param echo.Context
* @return error
 */
func (u UserController) Index(c echo.Context) (err error) {

	// Request Bind And Validation
	request := new(scopes.Pagination)

	if err = c.Bind(request); err != nil {
		return
	}

	var users []models.User
	users, err = u.UserService.GetUsers(request)
	if err != nil {
		return echo.ErrInternalServerError
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)

	// Policy Control
	if !u.UserPolicy.Index(claims) {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected "))
	}

	count, err := u.UserService.GetUsersCount()
	if err != nil {
		return echo.ErrInternalServerError
	}

	// Response
	return c.JSON(http.StatusOK, viewModels.SuccessResponse(viewModels.Paginator{
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

	// Request Bind And Validation

	request := new(requests.UserShowRequest)

	if err = c.Bind(request); err != nil {
		return
	}

	v := request.Validate()

	if v != nil {
		return c.JSON(http.StatusUnprocessableEntity, viewModels.ValidationResponse(v))
	}

	var user models.User
	user, err = u.UserService.GetUserByID(request.User)
	if err != nil {
		return echo.ErrInternalServerError
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)

	// Policy Control
	if !u.UserPolicy.Show(claims) {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected "))
	}

	// Response
	return c.JSON(http.StatusOK, viewModels.SuccessResponse(user))
}
