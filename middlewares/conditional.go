package GMiddleware

import (
	"github.com/labstack/echo/v4"
	"gotham/helpers"
	"net/http"
)

type MiddlewareI interface {
	control(c echo.Context) (bool, error)
}

// Conditional Middlewares
func Or(middleware []MiddlewareI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var msg = "You cannot access"

			for _, m := range middleware {
				canDo, err := m.control(c)
				if err != nil {
					msg = err.Error()
				}
				if canDo {
					return next(c)
				}
			}

			return c.JSON(http.StatusBadRequest, helpers.ErrorResponse(http.StatusBadRequest, msg))
		}
	}
}

func And(middleware []MiddlewareI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var msg = "You cannot access"

			for _, m := range middleware {
				canDo, err := m.control(c)
				if err != nil {
					msg = err.Error()
				}
				if !canDo {
					return c.JSON(http.StatusBadRequest, helpers.ErrorResponse(http.StatusBadRequest, msg))
				}
			}

			return next(c)
		}
	}
}

