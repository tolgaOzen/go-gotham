package controllers

import (
	"github.com/labstack/echo/v4"
	"gotham/viewModels"
	"net/http"
	"os"
)

type ServerController struct{}

/**
 * Ping
 *
 * @param echo.Context
 * @return error
 */
func (ServerController) Ping(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, viewModels.MResponse("pong"))
}

/**
 * Version
 *
 * @param echo.Context
 * @return error
 */
func (ServerController) Version(c echo.Context) (err error) {
	return c.JSON(http.StatusOK,map[string]interface{}{
		"version": os.Getenv("VERSION"),
	} )
}
