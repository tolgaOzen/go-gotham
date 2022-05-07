package GMiddleware

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"gotham/config"
	"gotham/services"
)

type Auth struct {
	UserService services.IUserService
}

func (s Auth) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*config.JwtCustomClaims)
		auth, err := s.UserService.GetUserByID(claims.AuthID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(401, "auth user could not be found")
			}
			return echo.ErrInternalServerError
		}
		c.Set("auth", auth)
		return next(c)
	}
}
