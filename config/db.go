package config

import "os"

type Database struct {
	DbConnection string
	DbDatabase   string
	DbHost       string
	DbPort       string
	DbUserName   string
	DbPassword   string
}

func GetDbConfig() Database {
	return Database{
		DbConnection: os.Getenv("DB_CONNECTION"),
		DbDatabase:   os.Getenv("DB_DATABASE"),
		DbHost:       os.Getenv("DB_HOST"),
		DbPort:       os.Getenv("DB_PORT"),
		DbUserName:   os.Getenv("DB_USERNAME"),
		DbPassword:   os.Getenv("DB_PASSWORD"),
	}
}
