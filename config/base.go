package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Conf *Config

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/**
 * Config
 *
 */
type Config struct {
	Port      string
	BaseUrl   string
	Db        Database
	SecretKey string
	Email     Email
	Brand     struct {
		ProjectName   string
		ProjectUrl    string
		ProjectApiUrl string
	}
}

/**
 * Configurations
 *
 */
func Configurations() {
	port := os.Getenv("API_PORT")
	Conf = &Config{
		Port:      port,
		BaseUrl:   os.Getenv("BASE_URL") + ":" + port,
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
		Email:     GetEmailConfig(),
		Brand: struct {
			ProjectName   string
			ProjectUrl    string
			ProjectApiUrl string
		}{ProjectName: os.Getenv("PROJECT_NAME"), ProjectUrl: os.Getenv("PROJECT_URL"), ProjectApiUrl: os.Getenv("PROJECT_API_URL")},
	}
}
