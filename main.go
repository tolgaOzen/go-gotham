package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/dingo/v4"
	"gotham/app"
	provider "gotham/app/provider"
	"gotham/config"
	"gotham/routers"
	"os"
)

func init() {
	production := flag.Bool("production", false, "a bool")
	flag.Parse()
	if !*production {
		err := dingo.GenerateContainer((*provider.Provider)(nil), "./app/container")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func main() {
	config.Configurations()
	app.New()
	defer app.Application.Container.Delete()

	//migrations.Initialize()
	//seeds.Initialize()

	routers.Route(echo.New())
}
