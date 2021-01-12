package dic

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	controllers "gotham/controllers"
	middlewares "gotham/middlewares"
	repositories "gotham/repositories"
	services "gotham/services"

	gorm "gorm.io/gorm"
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
					var eo *gorm.DB
					return eo, err
				}
				pi0, err := ctn.SafeGet("db-pool")
				if err != nil {
					var eo *gorm.DB
					return eo, err
				}
				p0, ok := pi0.(gorm.Dialector)
				if !ok {
					var eo *gorm.DB
					return eo, errors.New("could not cast parameter 0 to gorm.Dialector")
				}
				b, ok := d.Build.(func(gorm.Dialector) (*gorm.DB, error))
				if !ok {
					var eo *gorm.DB
					return eo, errors.New("could not cast build function to func(gorm.Dialector) (*gorm.DB, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				d, err := provider.Get("db")
				if err != nil {
					return err
				}
				c, ok := d.Close.(func(*gorm.DB) error)
				if !ok {
					return errors.New("could not cast close function to 'func(*gorm.DB) error'")
				}
				o, ok := obj.(*gorm.DB)
				if !ok {
					return errors.New("could not cast object to '*gorm.DB'")
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
					var eo gorm.Dialector
					return eo, err
				}
				b, ok := d.Build.(func() (gorm.Dialector, error))
				if !ok {
					var eo gorm.Dialector
					return eo, errors.New("could not cast build function to func() (gorm.Dialector, error)")
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
				b, ok := d.Build.(func(services.IUserService) (controllers.UserController, error))
				if !ok {
					var eo controllers.UserController
					return eo, errors.New("could not cast build function to func(services.IUserService) (controllers.UserController, error)")
				}
				return b(p0)
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
				p0, ok := pi0.(*gorm.DB)
				if !ok {
					var eo repositories.IUserRepository
					return eo, errors.New("could not cast parameter 0 to *gorm.DB")
				}
				b, ok := d.Build.(func(*gorm.DB) (repositories.IUserRepository, error))
				if !ok {
					var eo repositories.IUserRepository
					return eo, errors.New("could not cast build function to func(*gorm.DB) (repositories.IUserRepository, error)")
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
