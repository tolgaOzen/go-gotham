package controllers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gotham/config"
	"gotham/helpers"
	"gotham/requests"
	"gotham/services"
	"gotham/viewModels"
	"net/http"
	"time"
)

type AuthController struct {
	AuthService services.IAuthService
}

/**
 * Login
 *
 * @param echo.Context
 * @return error
 */
func (a AuthController) Login(c echo.Context) (err error) {

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

	user, dbError := a.AuthService.GetUserByEmail(request.Email)

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

	b, err := a.AuthService.Check(request.Email, request.Password)
	if !b {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"errors": map[string]string{
				"password": "password is incorrect",
			},
		})
	}

	accessTokenExp := time.Now().Add(time.Hour * 720).Unix()

	claims := &config.JwtCustomClaims{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(config.Conf.SecretKey))

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(viewModels.Login{
		AccessToken: accessToken,
		AccessTokenExp: accessTokenExp,
		User: user,
	}))
}
