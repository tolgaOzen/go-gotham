package main

import (
	"gotham/app"
	"gotham/config"
	"gotham/database/migrations"
	"gotham/database/seeds"
	"gotham/routers"
)

func main() {
	config.Configurations()
	app.New()
	defer app.Application.Container.Delete()
	migrations.Initialize()
	seeds.Initialize()
	routers.Initialize()
}
