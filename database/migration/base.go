package migrations

import (
	"gotham/app"
	"gotham/models"
)

func Initialize() {
	db := app.Application.Container.GetDb().DB()
	_ = db.AutoMigrate(&models.User{})
}
