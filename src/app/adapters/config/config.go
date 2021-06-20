package config

import (
	"flag"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port      int
	Auth      string
	JwtSecret string
	Services  map[string]Service
	Redis     redis.Options
}

type Service struct {
	Host    string
	Timeout time.Duration
}

func LoadConfig() Config {
	var cfg Config

	configPath := flag.String("config", "./config.yaml", "a string")
	flag.Parse()

	data, err := ioutil.ReadFile(*configPath)
	if err != nil {
		port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
		cfg.Port = port
		cfg.Auth = os.Getenv("AUTH_HOST")
		cfg.Services = make(map[string]Service)
		cfg.Services["position"] = Service{
			Host: os.Getenv("POSITION_HOST"),
			Timeout: 1000,
		}
		cfg.JwtSecret = os.Getenv("JWT_SECRET")
		cfg.Redis.Addr = os.Getenv("REDIS_HOST") + ":6379"
		cfg.Redis.Password = ""
		cfg.Redis.DB = 0

		return cfg
	}

	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		log.Fatalf("error: %v", err)
	}

	return cfg
}
