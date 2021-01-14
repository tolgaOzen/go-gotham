package GMiddleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gotham/config"
	"gotham/services"
)

type IsAdmin struct {
	UserService services.IUserService
}

func (i IsAdmin) control(c echo.Context) (bool bool, err error) {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*config.JwtCustomClaims)

	user, err := i.UserService.GetUserByID(int(claims.Id))

	if err != nil {
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
