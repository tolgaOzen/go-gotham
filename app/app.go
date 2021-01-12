package app

import (
	"github.com/sarulabs/di/v2"
	"gotham/app/container/dic"
	"log"
)

var Application *App

type App struct {
	Container *dic.Container
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

