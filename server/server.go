package server

import (
	"context"
	"gotham/config"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type IEchoServer interface {
	StartServer(e *echo.Echo)
}

type EchoServer struct {
	echo.Echo
	IEchoServer
}

func (ES *EchoServer) New() *echo.Echo {
	return echo.New()
}

func (ES *EchoServer) StartServer(e *echo.Echo) {
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
