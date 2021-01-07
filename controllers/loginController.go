package controllers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gotham/app"
	"gotham/app/container/dic"
	"gotham/config"
	"gotham/helpers"
	"gotham/models"
	"gotham/requests"
	"net/http"
	"time"
)

type LoginController struct{}

/**
 * Login
 *
 * @param echo.Context
 * @return error
 */
func (LoginController) Login(c echo.Context) (err error) {

	request := new(requests.LoginRequest)

	if err = c.Bind(request); err != nil {
		return
	}

	v := request.Validate()

	if v != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"errors": v,
		})
	}

	user := models.User{}
	dbError := dic.Db(c.Request()).First(&user, "email = ?", request.Email).Error

	if dbError != nil {
		if errors.Is(dbError, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
				"errors": map[string]string{
					"email": "email is incorrect",
				},
			})
		} else {
			return echo.ErrInternalServerError
		}
	}

	if !user.VerifyPassword(request.Password) {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"errors": map[string]string{
				"password": "password is incorrect",
			},
		})
	}

	exp := time.Now().Add(time.Minute * 15).Unix()

	claims := &config.JwtCustomClaims{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(app.Application.Config.SecretKey))
	if err != nil {
		return
	}

	data := map[string]interface{}{
		"access_token":      t,
		"access_token_exp":  exp,
		"user":              user,
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(data))
}
