package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Development      bool   `envconfig:"DEVELOPMENT" default:"false"`
	DatabaseURL      string `envconfig:"DATABASE_URL" default:"diploid.db"`
	BindAddress      string `default:":3000"`
	JWTSecret        string `envconfig:"JWT_SECRET"`
	CorsAllowOrigins string `envconfig:"CORS_ALLOW_ORIGINS" default:"http://localhost:3000,http://localhost:5173"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return &cfg
}
