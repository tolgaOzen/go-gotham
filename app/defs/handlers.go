package defs

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gotham/controllers"
	"gotham/services"
)

var HandlerDefs = []dingo.Def{
	{
		Name:  "user-controller",
		Scope: di.App,
		Build: func(service services.IUserService) (controllers.UserController, error) {
			return controllers.UserController{
				IUserService: service,
			}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-service"),
		},
	},
	{
		Name:  "auth-controller",
		Scope: di.App,
		Build: func(service services.IAuthService) (controllers.AuthController, error) {
			return controllers.AuthController{
				IAuthService: service,
			}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("auth-service"),
		},
	},
}



