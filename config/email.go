package config

import "os"

type EmailConfig struct {
	From string
	Host     string
	Port     string
	Password string
}

func GetEmailConfig() EmailConfig {
	return EmailConfig{
		From: os.Getenv("FROM"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Password: os.Getenv("PASSWORD"),
	}
}
