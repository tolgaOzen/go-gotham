package GMiddleware

import (
	"github.com/labstack/echo/v4"
)

type IsAdmin struct {
	UserId uint
}

func (i IsAdmin) control(c echo.Context) (bool bool, err error) {
	//u := c.Get("user").(*jwt.Token)
	//claims := u.Claims.(*config.JwtCustomClaims)

	//user := models.User{}
	//if err := app.Application.Container.GetDb().First(&user, claims.Id).Error; err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return false, echo.ErrUnauthorized
	//	}
	//	return false, echo.ErrInternalServerError
	//}

	//if user.IsAdmin() {
		return true, nil
	//}

	//return false, errors.New("you are not admin")
}
