package defs

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gotham/repositories"
	"gotham/services"
)

var ServicesDefs = []dingo.Def{
	{
		Name:  "auth-service",
		Scope: di.App,
		Build: func(repository repositories.IUserRepository) (s services.IAuthService, err error) {
			s = &services.AuthService{UserRepository: repository}
			return s, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-repository"),
		},
	},
	{
		Name:  "user-service",
		Scope: di.App,
		Build: func(repository repositories.IUserRepository) (s services.IUserService, err error) {
			return &services.UserService{UserRepository: repository}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-repository"),
		},
	},
}
