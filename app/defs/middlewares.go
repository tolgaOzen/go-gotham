package defs

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	GMiddleware "gotham/middlewares"
	"gotham/services"
)

var MiddlewaresDefs = []dingo.Def{
	{
		Name:  "is-admin-middleware",
		Scope: di.App,
		Build: func(repository services.IUserService) (s GMiddleware.IsAdmin, err error) {
			return GMiddleware.IsAdmin{UserService: repository}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-service"),
		},
	},
	{
		Name:  "is-verified-middleware",
		Scope: di.App,
		Build: func(repository services.IUserService) (s GMiddleware.IsVerified, err error) {
			return GMiddleware.IsVerified{UserService: repository}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-service"),
		},
	},
	{
		Name:  "auth-middleware",
		Scope: di.App,
		Build: func(repository services.IUserService) (s GMiddleware.Auth, err error) {
			return GMiddleware.Auth{UserService: repository}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-service"),
		},
	},
}
