package seeds

import "gotham/app"

func Initialize() {
	_ = app.Application.Container.GetUserRepository().Seed()
}

