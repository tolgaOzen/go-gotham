package controllers

import (
	"github.com/labstack/echo/v4"
	"gotham/helpers"
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
	return c.JSON(http.StatusOK, helpers.MResponse(200 , "pong"))
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
