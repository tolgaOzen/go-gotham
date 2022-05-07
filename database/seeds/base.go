package seeds

import (
	"gotham/app"
	"gotham/app/flags"
)

func Initialize() {
	if *flags.Seed {
		_ = app.Application.Container.GetUserRepository().Seed()
	}
}
