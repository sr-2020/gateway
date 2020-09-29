package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host      string
	JwtSecret string
	Login     string
	Password  string
	ModelId   int
}

func LoadConfig() Config {
	modelId, err := strconv.Atoi(os.Getenv("MODEL_ID"))
	if err != nil {
		log.Fatal(modelId)
	}

	return Config{
		Host:      os.Getenv("HOST"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		Login:     os.Getenv("LOGIN"),
		Password:  os.Getenv("PASSWORD"),
		ModelId:   modelId,
	}
}
