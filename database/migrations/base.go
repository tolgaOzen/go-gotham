package migrations

import (
	"gotham/app"
)

func Initialize() {
	_ = app.Application.Container.GetUserRepository().Migrate()
}
