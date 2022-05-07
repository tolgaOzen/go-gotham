package GMiddleware

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gotham/config"
	"gotham/services"
)

type IsVerified struct {
	UserService services.IUserService
}

func (i IsVerified) control(c echo.Context) *echo.HTTPError {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*config.JwtCustomClaims)

	user, err := i.UserService.GetUserByID(claims.AuthID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(404, "user could not be found")
		}
		return echo.ErrInternalServerError
	}

	if user.IsVerified() {
		return nil
	}

	return echo.NewHTTPError(403, "your email not verified")
}
