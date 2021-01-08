package defs

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gorm.io/gorm"
	"gotham/config"
	"gotham/services"
)

var DatabaseServiceDefs = []dingo.Def{
	{
		Name:  "db-pool",
		Scope: di.App,
		Build: func() (gorm.Dialector, error) {
			return services.NewDatabaseService(config.GetDbConfig()).OpenDatabase(), nil
		},
	},
	{
		Name:  "db",
		Scope: di.Request,
		Build: func(dia gorm.Dialector) (db *gorm.DB,err error) {
			return services.DatabaseService{}.ConnectDatabase(dia)
		},
		Params: dingo.Params{
			"0": dingo.Service("db-pool"),
		},
		Close: func(db *gorm.DB) error {
			sqlDB, _ := db.DB()
			return sqlDB.Close()
		},
	},
}
