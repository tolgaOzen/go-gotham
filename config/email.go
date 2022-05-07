package config

import "os"

type Email struct {
	From     string
	Host     string
	Port     string
	Password string
}

func GetEmailConfig() Email {
	return Email{
		From:     os.Getenv("FROM"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Password: os.Getenv("PASSWORD"),
	}
}
