package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	// Application
	Port      string
	SecretKey string

	// Database
	DatabaseURL string
}

func LoadEnvironment() (*Environment, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Environment{
		// Application
		Port:      getEnv("PORT", "8080"),
		SecretKey: os.Getenv("SECRET_KEY"),

		// Database
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
