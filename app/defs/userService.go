package defs

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gorm.io/gorm"
	"gotham/repositories"
	"gotham/services"
)

var UserServiceDefs = []dingo.Def{
	{
		Name:  "user-repository",
		Scope: di.App,
		Build: func(db *gorm.DB) (s repositories.IUserRepository, err error) {
			return &repositories.UserRepository{DB: db}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("db"),
		},
	},
	{
		Name:  "user-service",
		Scope: di.App,
		Build: func(repository repositories.IUserRepository) (s services.IUserService , err error) {
			return &services.UserService{UserRepository: repository}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-repository"),
		},
	},
}
