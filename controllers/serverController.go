package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"gotham/viewModels"
)

type ServerController struct{}

// Ping godoc
// @Tags Server
// @Success 200 {object} viewModels.Message{}
// @Failure 500
// @Router /status/ping [get]
func (ServerController) Ping(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, viewModels.MResponse("pong"))
}

// Version godoc
// @Tags Server
// @Success 200 {object} viewModels.Message{}
// @Failure 500
// @Router /status/version [get]
func (ServerController) Version(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"version": os.Getenv("VERSION"),
	})
}
