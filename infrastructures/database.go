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
 * GormDatabaseService
 *
 */
type GormDatabase struct {
	Pool     IGormDatabasePool
	Database *gorm.DB
}

/**
 * DB
 * get DB
 */
func (g *GormDatabase) DB() *gorm.DB {
	return g.Database
}

/**
 * NewGormDatabase
 *
 */
func NewGormDatabase(pool IGormDatabasePool) (*GormDatabase, error) {
	connection, err := gorm.Open(pool.GetDialector(), &gorm.Config{})
	return &GormDatabase{
		Pool:     pool,
		Database: connection,
	}, err
}

/**
 * IGormDatabasePool
 *
 */
type IGormDatabasePool interface {
	GetDialector() gorm.Dialector
}

/**
 * GormDatabasePool
 *
 */
type GormDatabasePool struct {
	Dialector gorm.Dialector
}

/**
 * GetDialector
 *
 */
func (m *GormDatabasePool) GetDialector() gorm.Dialector {
	return m.Dialector
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

/**
 * MysqlPool
 *
 */
type MysqlPool struct {
	GormDatabasePool
}

/**
 * NewMysqlPool
 *
 */
func NewMysqlPool(DbConfig config.Database) IGormDatabasePool {
	dsn := DbConfig.DbUserName + ":" + DbConfig.DbPassword + "@(" + DbConfig.DbHost + ")/" + DbConfig.DbDatabase + "?charset=utf8&parseTime=True&loc=Local"
	return &MysqlPool{
		GormDatabasePool{
			Dialector: mysql.Open(dsn),
		},
	}
}

/**
 * PostgresPool
 *
 */
type PostgresPool struct {
	GormDatabasePool
}

/**
 * NewPostgresPool
 *
 */
func NewPostgresPool(DbConfig config.Database) IGormDatabasePool {
	return &PostgresPool{
		GormDatabasePool{
			Dialector: postgres.New(postgres.Config{
				DSN:                  "user=" + DbConfig.DbUserName + " host=" + DbConfig.DbHost + " password=" + DbConfig.DbPassword + " dbname=" + DbConfig.DbDatabase + " port=" + DbConfig.DbPort + " sslmode=disable",
				PreferSimpleProtocol: true,
			}),
		},
	}
}
