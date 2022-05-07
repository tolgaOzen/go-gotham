package defs

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gotham/config"
	"gotham/infrastructures"
)

var InfrastructuresDefs = []dingo.Def{
	{
		Name:  "db-pool",
		Scope: di.App,
		Build: func() (infrastructures.IGormDatabasePool, error) {
			return infrastructures.NewGormDatabasePool(config.GetDbConfig()), nil
		},
		NotForAutoFill: true,
	},
	{
		Name:  "db",
		Scope: di.App,
		Build: func(pool infrastructures.IGormDatabasePool) (infrastructures.IGormDatabase, error) {
			return infrastructures.NewGormDatabase(pool)
		},
		Params: dingo.Params{
			"0": dingo.Service("db-pool"),
		},
		Close: func(db infrastructures.IGormDatabase) error {
			gormDB, _ := db.DB().DB()
			return gormDB.Close()
		},
	},
	{
		Name:  "email",
		Scope: di.App,
		Build: func() (emailService infrastructures.IEmailService, err error) {
			return infrastructures.NewEmailService(&config.Conf.Email), nil
		},
	},
}
