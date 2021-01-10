package GMiddleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/dingo/v4"
	"gotham/app"
)

// For dependency injection container's request scope.
func DicSubContainerSetterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctn, err := app.Application.Container.SubContainer()
		if err != nil {
			panic(err)
		}
		defer ctn.Delete()
		ctx := context.WithValue(c.Request().Context(), dingo.ContainerKey("dingo"), ctn)
		req := c.Request().WithContext(ctx)
		c.SetRequest(req)
		return next(c)
	}
}
