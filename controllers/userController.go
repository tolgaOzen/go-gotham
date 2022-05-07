package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gotham/models"
	"gotham/policies"
	"gotham/requests"
	"gotham/services"
	"gotham/viewModels"
)

type UserController struct {
	UserService services.IUserService

	UserPolicy policies.IUserPolicy
}

// Index godoc
// @Summary List of users
// @Description
// @Tags User
// @Accept  json
// @Accept  multipart/form-data
// @Accept  application/x-www-form-urlencoded
// @Produce json
// @Param token header string true "Bearer Token"
// @Success 200 {object} viewModels.Paginator{data=[]models.User}
// @Failure 400 {object} viewModels.Message{}
// @Failure 401 {object} viewModels.Message{}
// @Failure 500 {object} viewModels.Message{}
// @Router /v1/r/users [get]
func (u UserController) Index(c echo.Context) (err error) {
	auth := models.ConvertUser(c.Get("auth"))

	// Request Bind And Validation
	request := new(requests.UserIndexRequest)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &request.QueryParams); err != nil {
		return err
	}

	// Policy Control
	if !u.UserPolicy.Index(auth) {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected "))
	}

	var count int64
	var users []models.User
	users, count, err = u.UserService.GetUsersWithPaginationAndOrder(&request.QueryParams.Pagination, &request.QueryParams.Order)
	if err != nil {
		return echo.ErrInternalServerError
	}

	// Response
	return c.JSON(http.StatusOK, viewModels.SuccessResponse(viewModels.Paginator{
		TotalRecord: count,
		Records:     users,
		Limit:       request.QueryParams.Pagination.GetLimit(),
		Page:        request.QueryParams.Pagination.GetPage(),
	}))
}

// Show godoc
// @Summary Get User
// @Description
// @Tags User
// @Accept  json
// @Accept  multipart/form-data
// @Accept  application/x-www-form-urlencoded
// @Produce json
// @Param token header string true "Bearer Token"
// @Success 200 {object} viewModels.HTTPSuccessResponse{data=models.User}
// @Failure 404 {object} viewModels.Message{}
// @Failure 401 {object} viewModels.Message{}
// @Failure 400 {object} viewModels.Message{}
// @Failure 403 {object} viewModels.Message{}
// @Failure 500 {object} viewModels.Message{}
// @Router /v1/r/users/:user [get]
func (u UserController) Show(c echo.Context) (err error) {
	auth := models.ConvertUser(c.Get("auth"))

	// Request Bind And Validation

	request := new(requests.UserShowRequest)

	if err := (&echo.DefaultBinder{}).BindPathParams(c, &request.PathParams); err != nil {
		return err
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &request.Body); err != nil {
		return err
	}

	v := request.Validate()
	if v != nil {
		return c.JSON(http.StatusUnprocessableEntity, viewModels.ValidationResponse(v))
	}

	var user models.User
	user, err = u.UserService.GetUserByID(request.PathParams.User)
	if err != nil {
		return echo.ErrInternalServerError
	}

	// Policy Control
	if !u.UserPolicy.Show(auth, user) {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected "))
	}

	// Response
	return c.JSON(http.StatusOK, viewModels.SuccessResponse(user))
}
