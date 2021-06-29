package routers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
	"os/signal"
	"time"

	"gotham/app"
	"gotham/config"
	"gotham/controllers"
	"gotham/docs"
	GMiddleware "gotham/middlewares"
)



func Route(e *echo.Echo) {

	docs.SwaggerInfo.Title = "Gotham API"
	docs.SwaggerInfo.Description = "..."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "''"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"v1"}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/doc/*", echoSwagger.WrapHandler)

	//server
	e.GET("/status/ping", controllers.ServerController{}.Ping)
	e.GET("/status/version", controllers.ServerController{}.Version)

	v1 := e.Group("/v1")

	//login
	v1.POST("/login", app.Application.Container.GetAuthController().Login)

	r := v1.Group("/restricted")

	c := middleware.JWTConfig{
		Claims:     &config.JwtCustomClaims{},
		SigningKey: []byte(config.Conf.SecretKey),
	}

	r.Use(middleware.JWTWithConfig(c))
	r.Use(app.Application.Container.GetAuthMiddleware().AuthMiddleware)

	//user
	r.GET("/users/:user", app.Application.Container.GetUserController().Show,GMiddleware.Or(app.Application.Container.GetIsAdminMiddleware(), app.Application.Container.GetIsVerifiedMiddleware()))
	r.GET("/users", app.Application.Container.GetUserController().Index)


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
