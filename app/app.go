package app

import (
	"fmt"
	"log"
	"os"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	"gotham/app/container/dic"
	"gotham/app/flags"
	"gotham/app/provider"
)

var Application *App

type App struct {
	Container *dic.Container
}

func init() {
	if !*flags.Production {
		err := dingo.GenerateContainer((*provider.Provider)(nil), "./app/container")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

/**
 * New
 *
 */
func New() {
	Application = &App{}
	container, err := dic.NewContainer(di.App)
	if err != nil {
		log.Fatal("Error dic.NewContainer")
	}
	Application.Container = container
}
