package app

import (
	"github.com/joho/godotenv"
	"github.com/sarulabs/di/v2"
	"gotham/app/container/dic"
	"gotham/config"
	"log"
	"os"
)

var Application *App

type App struct {
	Container *dic.Container
	Config    *Config
}

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/**
 * New
 *
 */
func New() {
	Application = &App{}
	Application.Config = Configurations()
	container, err := dic.NewContainer(di.App, di.Request)
	if err != nil {
		log.Fatal("Error dic.NewContainer")
	}
	Application.Container = container
}

/**
 * Config
 *
 */
type Config struct {
	Port      string
	BaseUrl   string
	Db        config.Database
	SecretKey string
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
func Configurations() *Config {
	port := os.Getenv("API_PORT")
	return &Config{
		Port:      port,
		BaseUrl:   os.Getenv("BASE_URL") + ":" + port,
		//Db:        config.GetDbConfig(),
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
		Brand: struct {
			ProjectName   string
			ProjectUrl    string
			ProjectApiUrl string
		}{ProjectName: os.Getenv("PROJECT_NAME"), ProjectUrl: os.Getenv("PROJECT_URL"), ProjectApiUrl: os.Getenv("PROJECT_API_URL")},
	}
}
