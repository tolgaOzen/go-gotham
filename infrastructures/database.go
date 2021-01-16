package infrastructures

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotham/config"
)

/**
 * IGormDatabase
 *
 */
type IGormDatabase interface {
	DB() *gorm.DB
}

/**
 * IGormDatabasePool
 *
 */
type IGormDatabasePool interface {
	GetDialector() gorm.Dialector
}

/**
 * GormDatabaseService
 *
 */
type GormDatabaseService struct {
	Pool     IGormDatabasePool
	Database *gorm.DB
}

/**
 * NewGormDatabasePool
 *
 */
func NewGormDatabasePool(dbConfig config.Database) IGormDatabasePool {
	switch dbConfig.DbConnection {
	case "postgres":
		return NewPostgresPool(dbConfig)
	case "mysql":
		return NewMysqlPool(dbConfig)
	default:
		return NewMysqlPool(dbConfig)
	}
}

func (g *GormDatabaseService) DB() *gorm.DB {
	return g.Database
}

func NewGormDatabase(pool IGormDatabasePool) (IGormDatabase, error) {
	connection ,err := gorm.Open(pool.GetDialector(), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	var db = &GormDatabaseService{
		Pool: pool,
		Database: connection,
	}
	return db, err
}


/**
 * MysqlPool
 *
 */
type MysqlPool struct{
	Dialector gorm.Dialector
}

func NewMysqlPool(DbConfig config.Database) MysqlPool {
	dsn := DbConfig.DbUserName + ":" + DbConfig.DbPassword + "@(" + DbConfig.DbHost + ")/" + DbConfig.DbDatabase + "?charset=utf8&parseTime=True&loc=Local"
	return MysqlPool{
		Dialector: mysql.Open(dsn),
	}
}

func (m MysqlPool) GetDialector() gorm.Dialector {
	return m.Dialector
}


/**
 * PostgresPool
 *
 */
type PostgresPool struct{
	Dialector gorm.Dialector
}

func NewPostgresPool(DbConfig config.Database) PostgresPool {
	return PostgresPool{
		Dialector: postgres.New(postgres.Config{
			DSN:                  "user=" + DbConfig.DbUserName + " host=" + DbConfig.DbHost + " password=" + DbConfig.DbPassword + " dbname=" + DbConfig.DbDatabase + " port=" + DbConfig.DbPort + " sslmode=disable",
			PreferSimpleProtocol: true,
		}),
	}
}

func (m PostgresPool) GetDialector() gorm.Dialector {
	return m.Dialector
}

