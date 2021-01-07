package dic

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	gorm "gorm.io/gorm"
)

func getDiDefs(provider dingo.Provider) []di.Def {
	return []di.Def{
		{
			Name:  "db",
			Scope: "request",
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
	}
}
