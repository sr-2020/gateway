package config

import (
	"os"
)

type Config struct {
	Host      string
	JwtSecret string
}

func LoadConfig() Config {
	return Config{
		Host: os.Getenv("HOST"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
