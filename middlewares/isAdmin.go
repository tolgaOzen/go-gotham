package GMiddleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gotham/app/container/dic"
	"gotham/config"
	"gotham/models"
)

type IsAdmin struct {}

func (i IsAdmin) control(c echo.Context) (bool bool, err error) {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*config.JwtCustomClaims)

	user := models.User{}
	if err := dic.Db(c.Request()).First(&user, claims.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, echo.ErrUnauthorized
		}
		return false, echo.ErrInternalServerError
	}


	if user.IsAdmin() {
		return true, nil
	}

	return false, errors.New("you are not admin")
}
