package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gotham/app"
	"gotham/config"
	"gotham/controllers"
	GMiddleware "gotham/middlewares"
	"net/http"
)



func Route(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	//e.Use(GMiddleware.DicSubContainerSetterMiddleware)

	//server
	e.GET("/status/ping", controllers.ServerController{}.Ping)
	e.GET("/status/version", controllers.ServerController{}.Version)

	//login
	e.POST("/login", app.Application.Container.GetAuthController().Login)

	r := e.Group("/restricted")

	c := middleware.JWTConfig{
		Claims:     &config.JwtCustomClaims{},
		SigningKey: []byte(config.Conf.SecretKey),
	}

	r.Use(middleware.JWTWithConfig(c))

	//user
	r.GET("/users/:user", app.Application.Container.GetUserController().Show,GMiddleware.Or([]GMiddleware.IMiddleware{app.Application.Container.GetIsVerifiedMiddleware()}))
	r.GET("/users", app.Application.Container.GetUserController().Index,GMiddleware.And([]GMiddleware.IMiddleware{app.Application.Container.GetIsAdminMiddleware(),app.Application.Container.GetIsVerifiedMiddleware()}))

	s := &http.Server{
		Addr: ":" + config.Conf.Port,
	}

	e.Logger.Fatal(e.StartServer(s))
}
