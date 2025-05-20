package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	Email    string
	Password string
	Address  string
}

type Config struct {
	EmailConfig EmailConfig
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file, using default config")
	}
	return &Config{
		EmailConfig: EmailConfig{
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address:  os.Getenv("ADDRESS"),
		},
	}
}
