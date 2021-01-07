package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gotham/app"
	"gotham/controllers"
	GMiddleware "gotham/middlewares"
	"net/http"
)

func Route(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//server
	e.GET("/status/ping", controllers.ServerController{}.Ping)
	e.GET("/status/version", controllers.ServerController{}.Version)

	//user
	e.GET("/users/:user", controllers.UserController{}.Show, GMiddleware.Or([]GMiddleware.MiddlewareI{GMiddleware.IsAdmin{}, GMiddleware.IsVerified{}}))
	e.GET("/users", controllers.UserController{}.Index, GMiddleware.And([]GMiddleware.MiddlewareI{GMiddleware.IsVerified{}}))

	s := &http.Server{
		Addr: ":" + app.Application.Config.Port,
	}

	e.Logger.Fatal(e.StartServer(s))
}
