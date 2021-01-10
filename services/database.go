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
	open() gorm.Dialector
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
		d = Postgres{s}
	case "mysql":
		d = Mysql{s}
	default:
		d = Mysql{s}
	}
	db = d.open()
	return
}

func (DatabaseService) ConnectDatabase(dialector gorm.Dialector) (db *gorm.DB, err error) {
	return gorm.Open(dialector, &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
}

// Mysql
type Mysql struct{
	DatabaseService
}

func (m Mysql) open() (dia gorm.Dialector) {
	dsn := m.DbConfig.DbUserName + ":" + m.DbConfig.DbPassword + "@(" + m.DbConfig.DbHost + ")/" + m.DbConfig.DbDatabase + "?charset=utf8&parseTime=True&loc=Local"
	return mysql.Open(dsn)
}

// Postgresql
type Postgres struct{
	DatabaseService
}

func (p Postgres) open() (dia gorm.Dialector) {
	return postgres.New(postgres.Config{
		DSN:                  "user=" + p.DbConfig.DbUserName + " host=" + p.DbConfig.DbHost + " password=" + p.DbConfig.DbPassword + " dbname=" + p.DbConfig.DbDatabase + " port=" + p.DbConfig.DbPort + " sslmode=disable",
		PreferSimpleProtocol: true,
	})
}
