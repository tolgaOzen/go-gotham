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
	e.Use(GMiddleware.DicMiddleware)

	//server
	e.GET("/status/ping", controllers.ServerController{}.Ping)
	e.GET("/status/version", controllers.ServerController{}.Version)

	//login
	e.POST("/login", controllers.LoginController{}.Login)

	r := e.Group("/restricted")

	c := middleware.JWTConfig{
		Claims:     &config.JwtCustomClaims{},
		SigningKey: []byte(app.Application.Config.SecretKey),
	}

	r.Use(middleware.JWTWithConfig(c))

	//user
	r.GET("/users/:user", controllers.UserController{}.Show, GMiddleware.Or([]GMiddleware.MiddlewareI{GMiddleware.IsAdmin{}, GMiddleware.IsVerified{}}))
	r.GET("/users", controllers.UserController{}.Index, GMiddleware.And([]GMiddleware.MiddlewareI{GMiddleware.IsVerified{}}))

	s := &http.Server{
		Addr: ":" + app.Application.Config.Port,
	}


	e.Logger.Fatal(e.StartServer(s))
}
