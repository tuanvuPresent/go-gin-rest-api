package config

import (
	"os"
)

type Conf struct {
	DB_CONFIG string
	APP_PORT  string
}

var Config = Conf{}

func Init() {
	Config = Conf{
		DB_CONFIG: os.Getenv("DB_CONFIG"),
		APP_PORT:  os.Getenv("APP_PORT"),
	}
}
