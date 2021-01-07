package services

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotham/config"
	"sync"
)

var DatabaseServiceDefs = []dingo.Def{
	{
		Name:  "db-pool",
		Scope: di.App,
		Build: func() (gorm.Dialector, error) {
			var d = config.GetDbConfig()
			return Open(&d), nil
		},
	},
	{
		Name:  "db",
		Scope: di.App,
		Build: func(dia gorm.Dialector) (db *gorm.DB,err error) {
			return ConnectDatabase(dia)
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


type DatabaseConnecter interface {
	open(*config.Database) gorm.Dialector
}

var once sync.Once

func Open(dbConfig *config.Database) (db gorm.Dialector) {

	var d DatabaseConnecter

	switch dbConfig.DbConnection {
	case "postgres":
		d = Postgres{}
	case "mysql":
		d = Mysql{}
	default:
		d = Mysql{}
	}

	db = d.open(dbConfig)

	return
}

type Mysql struct{}

func (Mysql) open(dbConfig *config.Database) (dia gorm.Dialector) {
	dsn := dbConfig.DbUserName + ":" + dbConfig.DbPassword + "@(" + dbConfig.DbHost + ")/" + dbConfig.DbDatabase + "?charset=utf8&parseTime=True&loc=Local"
	return mysql.Open(dsn)
}


type Postgres struct{}

func (Postgres) open(dbConfig *config.Database) (dia gorm.Dialector) {
	return postgres.New(postgres.Config{
		DSN:  "user=" + dbConfig.DbUserName + " host=" + dbConfig.DbHost + " password=" + dbConfig.DbPassword + " dbname=" + dbConfig.DbDatabase + " port=" + dbConfig.DbPort + " sslmode=disable",
		PreferSimpleProtocol: true,
	})
}

func ConnectDatabase(dialector gorm.Dialector) (db *gorm.DB, err error) {
	once.Do(func() {
		db, err = gorm.Open(dialector, &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	})
	return
}
