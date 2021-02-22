package routers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gotham/app"
	"gotham/config"
	"gotham/controllers"
	GMiddleware "gotham/middlewares"
	"os"
	"os/signal"
	"time"
)



func Route(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

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
	r.GET("/users/:user", app.Application.Container.GetUserController().Show,GMiddleware.Or(app.Application.Container.GetIsVerifiedMiddleware()))
	r.GET("/users", app.Application.Container.GetUserController().Index,GMiddleware.And(app.Application.Container.GetIsAdminMiddleware(),app.Application.Container.GetIsVerifiedMiddleware()))


	// Start server
	go func() {
		if err := e.Start(":" + config.Conf.Port); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
