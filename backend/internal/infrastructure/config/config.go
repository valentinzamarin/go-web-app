package config

import "os"

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	/*
		func without err
		cuz db url sure will be
	*/
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
