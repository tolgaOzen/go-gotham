package services

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotham/config"
)

type DatabaseService struct {
	DbConfig config.Database
}

type DatabaseConnecter interface {
	open(config.Database) gorm.Dialector
}

func NewDatabaseService(dbConfig config.Database) *DatabaseService {
	return &DatabaseService{
		DbConfig: dbConfig,
	}
}

func (s DatabaseService) OpenDatabase() (db gorm.Dialector) {
	var d DatabaseConnecter
	switch s.DbConfig.DbConnection {
	case "postgres":
		d = Postgres{}
	case "mysql":
		d = Mysql{}
	default:
		d = Mysql{}
	}
	db = d.open(s.DbConfig)
	return
}

func (DatabaseService) ConnectDatabase(dialector gorm.Dialector) (db *gorm.DB, err error) {
	return gorm.Open(dialector, &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
}

type Mysql struct{}

func (Mysql) open(dbConfig config.Database) (dia gorm.Dialector) {
	dsn := dbConfig.DbUserName + ":" + dbConfig.DbPassword + "@(" + dbConfig.DbHost + ")/" + dbConfig.DbDatabase + "?charset=utf8&parseTime=True&loc=Local"
	return mysql.Open(dsn)
}

type Postgres struct{}

func (Postgres) open(dbConfig config.Database) (dia gorm.Dialector) {
	return postgres.New(postgres.Config{
		DSN:                  "user=" + dbConfig.DbUserName + " host=" + dbConfig.DbHost + " password=" + dbConfig.DbPassword + " dbname=" + dbConfig.DbDatabase + " port=" + dbConfig.DbPort + " sslmode=disable",
		PreferSimpleProtocol: true,
	})
}
