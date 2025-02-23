package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Cfg *Config

type Config struct {
	Environment string `envconfig:"ENV" default:"dev"`

	PageLimit int `envconfig:"PAGE_LIMIT" default:"10"`

	Server struct {
		Host         string        `envconfig:"SERVER_HOST" default:"0.0.0.0"`
		Port         int           `envconfig:"SERVER_PORT" default:"8080"`
		ReadTimeout  time.Duration `envconfig:"SERVER_READ_TIMEOUT" default:"30s"`
		WriteTimeout time.Duration `envconfig:"SERVER_WRITE_TIMEOUT" default:"30s"`
	}

	Database struct {
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Port     int    `envconfig:"DB_PORT" default:"5432"`
		User     string `envconfig:"DB_USER" required:"true"`
		Password string `envconfig:"DB_PASSWORD" required:"true"`
		Name     string `envconfig:"DB_NAME" default:"postgres"`
		SSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
		PoolSize int    `envconfig:"DB_POOL_SIZE" default:"10"`
	}

	Redis struct {
		Host     string `envconfig:"REDIS_HOST" default:"localhost"`
		Port     int    `envconfig:"REDIS_PORT" default:"6379"`
		Password string `envconfig:"REDIS_PASSWORD"`
		DB       int    `envconfig:"REDIS_DB" default:"0"`
	}
}

func LoadENV() {
	_ = godotenv.Load()

	Cfg = &Config{}

	if err := envconfig.Process("", Cfg); err != nil {
		fmt.Printf("config load error: %v", err)
	}
}
