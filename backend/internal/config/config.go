package config

import "os"

type Config struct {
	AppPort string

	DBHost string
	DBPort string

	DBName     string
	DBUser     string
	DBPassword string
}

func Load() Config {
	return Config{
		AppPort: os.Getenv("APP_PORT"),

		DBHost: os.Getenv("POSTGRES_HOST"),
		DBPort: os.Getenv("POSTGRES_PORT"),

		DBName:     os.Getenv("POSTGRES_DB"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
	}
}
