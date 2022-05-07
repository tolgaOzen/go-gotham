package GMiddleware

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"gotham/viewModels"
)

type IConditionalMiddleware interface {
	control(c echo.Context) *echo.HTTPError
}

// Conditional Middlewares

func Or(middleware ...IConditionalMiddleware) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err *echo.HTTPError
			for _, m := range middleware {
				err = m.control(c)
				if err == nil {
					return next(c)
				}
			}

			if err != nil {
				return c.JSON(err.Code, viewModels.MResponse(fmt.Sprintf("%v", err.Message)))
			}

			return next(c)
		}
	}
}

func And(middleware ...IConditionalMiddleware) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, m := range middleware {
				err := m.control(c)
				if err != nil {
					return c.JSON(err.Code, viewModels.MResponse(fmt.Sprintf("%v", err.Message)))
				}
			}
			return next(c)
		}
	}
}
