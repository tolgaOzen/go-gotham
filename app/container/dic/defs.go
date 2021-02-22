package dic

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	controllers "gotham/controllers"
	infrastructures "gotham/infrastructures"
	middlewares "gotham/middlewares"
	policies "gotham/policies"
	repositories "gotham/repositories"
	services "gotham/services"
)

func getDiDefs(provider dingo.Provider) []di.Def {
	return []di.Def{
		{
			Name:  "auth-controller",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("auth-controller")
				if err != nil {
					var eo controllers.AuthController
					return eo, err
				}
				pi0, err := ctn.SafeGet("auth-service")
				if err != nil {
					var eo controllers.AuthController
					return eo, err
				}
				p0, ok := pi0.(services.IAuthService)
				if !ok {
					var eo controllers.AuthController
					return eo, errors.New("could not cast parameter 0 to services.IAuthService")
				}
				b, ok := d.Build.(func(services.IAuthService) (controllers.AuthController, error))
				if !ok {
					var eo controllers.AuthController
					return eo, errors.New("could not cast build function to func(services.IAuthService) (controllers.AuthController, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
		{
			Name:  "auth-service",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("auth-service")
				if err != nil {
					var eo services.IAuthService
					return eo, err
				}
				pi0, err := ctn.SafeGet("user-repository")
				if err != nil {
					var eo services.IAuthService
					return eo, err
				}
				p0, ok := pi0.(repositories.IUserRepository)
				if !ok {
					var eo services.IAuthService
					return eo, errors.New("could not cast parameter 0 to repositories.IUserRepository")
				}
				b, ok := d.Build.(func(repositories.IUserRepository) (services.IAuthService, error))
				if !ok {
					var eo services.IAuthService
					return eo, errors.New("could not cast build function to func(repositories.IUserRepository) (services.IAuthService, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
		{
			Name:  "db",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("db")
				if err != nil {
					var eo infrastructures.IGormDatabase
					return eo, err
				}
				pi0, err := ctn.SafeGet("db-pool")
				if err != nil {
					var eo infrastructures.IGormDatabase
					return eo, err
				}
				p0, ok := pi0.(infrastructures.IGormDatabasePool)
				if !ok {
					var eo infrastructures.IGormDatabase
					return eo, errors.New("could not cast parameter 0 to infrastructures.IGormDatabasePool")
				}
				b, ok := d.Build.(func(infrastructures.IGormDatabasePool) (infrastructures.IGormDatabase, error))
				if !ok {
					var eo infrastructures.IGormDatabase
					return eo, errors.New("could not cast build function to func(infrastructures.IGormDatabasePool) (infrastructures.IGormDatabase, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				d, err := provider.Get("db")
				if err != nil {
					return err
				}
				c, ok := d.Close.(func(infrastructures.IGormDatabase) error)
				if !ok {
					return errors.New("could not cast close function to 'func(infrastructures.IGormDatabase) error'")
				}
				o, ok := obj.(infrastructures.IGormDatabase)
				if !ok {
					return errors.New("could not cast object to 'infrastructures.IGormDatabase'")
				}
				return c(o)
			},
		},
		{
			Name:  "db-pool",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("db-pool")
				if err != nil {
					var eo infrastructures.IGormDatabasePool
					return eo, err
				}
				b, ok := d.Build.(func() (infrastructures.IGormDatabasePool, error))
				if !ok {
					var eo infrastructures.IGormDatabasePool
					return eo, errors.New("could not cast build function to func() (infrastructures.IGormDatabasePool, error)")
				}
				return b()
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
		{
			Name:  "is-admin-middleware",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("is-admin-middleware")
				if err != nil {
					var eo middlewares.IsAdmin
					return eo, err
				}
				pi0, err := ctn.SafeGet("user-service")
				if err != nil {
					var eo middlewares.IsAdmin
					return eo, err
				}
				p0, ok := pi0.(services.IUserService)
				if !ok {
					var eo middlewares.IsAdmin
					return eo, errors.New("could not cast parameter 0 to services.IUserService")
				}
				b, ok := d.Build.(func(services.IUserService) (middlewares.IsAdmin, error))
				if !ok {
					var eo middlewares.IsAdmin
					return eo, errors.New("could not cast build function to func(services.IUserService) (middlewares.IsAdmin, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
		{
			Name:  "is-verified-middleware",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("is-verified-middleware")
				if err != nil {
					var eo middlewares.IsVerified
					return eo, err
				}
				pi0, err := ctn.SafeGet("user-service")
				if err != nil {
					var eo middlewares.IsVerified
					return eo, err
				}
				p0, ok := pi0.(services.IUserService)
				if !ok {
					var eo middlewares.IsVerified
					return eo, errors.New("could not cast parameter 0 to services.IUserService")
				}
				b, ok := d.Build.(func(services.IUserService) (middlewares.IsVerified, error))
				if !ok {
					var eo middlewares.IsVerified
					return eo, errors.New("could not cast build function to func(services.IUserService) (middlewares.IsVerified, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
		{
			Name:  "user-controller",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("user-controller")
				if err != nil {
					var eo controllers.UserController
					return eo, err
				}
				pi0, err := ctn.SafeGet("user-service")
				if err != nil {
					var eo controllers.UserController
					return eo, err
				}
				p0, ok := pi0.(services.IUserService)
				if !ok {
					var eo controllers.UserController
					return eo, errors.New("could not cast parameter 0 to services.IUserService")
				}
				pi1, err := ctn.SafeGet("user-policy")
				if err != nil {
					var eo controllers.UserController
					return eo, err
				}
				p1, ok := pi1.(policies.IUserPolicy)
				if !ok {
					var eo controllers.UserController
					return eo, errors.New("could not cast parameter 1 to policies.IUserPolicy")
				}
				b, ok := d.Build.(func(services.IUserService, policies.IUserPolicy) (controllers.UserController, error))
				if !ok {
					var eo controllers.UserController
					return eo, errors.New("could not cast build function to func(services.IUserService, policies.IUserPolicy) (controllers.UserController, error)")
				}
				return b(p0, p1)
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
		{
			Name:  "user-policy",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("user-policy")
				if err != nil {
					var eo policies.IUserPolicy
					return eo, err
				}
				b, ok := d.Build.(func() (policies.IUserPolicy, error))
				if !ok {
					var eo policies.IUserPolicy
					return eo, errors.New("could not cast build function to func() (policies.IUserPolicy, error)")
				}
				return b()
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
		{
			Name:  "user-repository",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("user-repository")
				if err != nil {
					var eo repositories.IUserRepository
					return eo, err
				}
				pi0, err := ctn.SafeGet("db")
				if err != nil {
					var eo repositories.IUserRepository
					return eo, err
				}
				p0, ok := pi0.(infrastructures.IGormDatabase)
				if !ok {
					var eo repositories.IUserRepository
					return eo, errors.New("could not cast parameter 0 to infrastructures.IGormDatabase")
				}
				b, ok := d.Build.(func(infrastructures.IGormDatabase) (repositories.IUserRepository, error))
				if !ok {
					var eo repositories.IUserRepository
					return eo, errors.New("could not cast build function to func(infrastructures.IGormDatabase) (repositories.IUserRepository, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
		{
			Name:  "user-service",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("user-service")
				if err != nil {
					var eo services.IUserService
					return eo, err
				}
				pi0, err := ctn.SafeGet("user-repository")
				if err != nil {
					var eo services.IUserService
					return eo, err
				}
				p0, ok := pi0.(repositories.IUserRepository)
				if !ok {
					var eo services.IUserService
					return eo, errors.New("could not cast parameter 0 to repositories.IUserRepository")
				}
				b, ok := d.Build.(func(repositories.IUserRepository) (services.IUserService, error))
				if !ok {
					var eo services.IUserService
					return eo, errors.New("could not cast build function to func(repositories.IUserRepository) (services.IUserService, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				return nil
			},
		},
	}
}
