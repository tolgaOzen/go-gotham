package migrations

import (
	"gotham/app"
	"gotham/app/flags"
)

func Initialize() {
	if *flags.Migrate {
		_ = app.Application.Container.GetUserRepository().Migrate()
	}
}
