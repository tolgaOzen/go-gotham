package routers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sarulabs/dingo/v4"
	"gotham/app"
	"gotham/controllers"
	GMiddleware "gotham/middlewares"
	"net/http"
)

// For dependency injection container's request scope.
func DicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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


func Route(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(DicMiddleware)

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
