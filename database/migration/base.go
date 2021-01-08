package migrations

import (
	"gotham/app"
	"gotham/models"
)

func Initialize() {
	db := app.Application.Container.UnscopedGetDb()

	_ = db.AutoMigrate(&models.User{})

	app.Application.Container.Clean()
}
