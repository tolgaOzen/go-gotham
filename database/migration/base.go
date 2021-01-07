package migrations

import (
	"gotham/app"
	"gotham/models"
)

func Initialize() {
	_ = app.Application.Container.GetDb().AutoMigrate(&models.User{})
}
