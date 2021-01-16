package defs

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gotham/infrastructures"
	"gotham/repositories"
)

var RepositoriesDefs = []dingo.Def{
	{
		Name:  "user-repository",
		Scope: di.App,
		Build: func(gormDatabase infrastructures.IGormDatabase) (s repositories.IUserRepository, err error) {
			return &repositories.UserRepository{IGormDatabase: gormDatabase}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("db"),
		},
	},
}
