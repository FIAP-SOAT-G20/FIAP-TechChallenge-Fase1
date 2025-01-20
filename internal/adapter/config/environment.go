package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	// Application
	Port      string
	SecretKey string
	Duration  string

	// Database
	DatabaseURL    string
	AppEnvironment string

	// Mercado Pago
	PaymentGatewayToken           string
	PaymentGatewayNotificationURL string
	PaymentGatewayURL             string
}

func LoadEnvironment() (*Environment, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return &Environment{
		// Application
		Port:           getEnv("PORT", "8080"),
		SecretKey:      os.Getenv("SECRET_KEY"),
		AppEnvironment: getEnv("APP_ENV", "development"),
		Duration:       getEnv("TOKEN_DURATION", "1h"),

		// Database
		DatabaseURL: os.Getenv("DATABASE_URL"),

		// Mercado Pago
		PaymentGatewayToken:           os.Getenv("MERCADO_PAGO_TOKEN"),
		PaymentGatewayNotificationURL: os.Getenv("MERCADO_PAGO_NOTIFICATION_URL"),
		PaymentGatewayURL:             os.Getenv("MERCADO_PAGO_URL"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
