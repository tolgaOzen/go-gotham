package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/dingo/v4"
	"gotham/app"
	provider "gotham/app/provider"
	"gotham/config"
	"gotham/database/migrations"
	"gotham/routers"
	"os"
)

func init() {
	err := dingo.GenerateContainer((*provider.Provider)(nil), "./app/container")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	config.Configurations()
	app.New()
	defer app.Application.Container.Delete()

	migrations.Initialize()
	//procedures.Initialize()
	//go jobs.Initialize()

	routers.Route(echo.New())
}
