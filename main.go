package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/dingo/v4"
	"gotham/app"
	provider "gotham/app/provider"
	migrations "gotham/database/migration"
	"gotham/models/procedures"
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
	app.New()
	defer app.Application.Container.Delete()

	migrations.Initialize()
	procedures.Initialize()
	//go jobs.Initialize()

	routers.Route(echo.New())
}
