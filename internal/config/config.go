package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                    string
	DatabaseURL             string
	FirebaseProjectID       string
	FirebaseCredentialsFile string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	config := Config{
		Port:                    getEnv("PORT", "8080"),
		DatabaseURL:             os.Getenv("DATABASE_URL"),
		FirebaseProjectID:       os.Getenv("FIREBASE_PROJECT_ID"),
		FirebaseCredentialsFile: os.Getenv("FIREBASE_CREDENTIALS_FILE"),
	}

	if config.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	if config.FirebaseProjectID == "" {
		return Config{}, fmt.Errorf("FIREBASE_PROJECT_ID is required")
	}
	if config.FirebaseCredentialsFile == "" {
		return Config{}, fmt.Errorf("FIREBASE_CREDENTIALS_FILE is required")
	}

	credentialsPath, err := filepath.Abs(config.FirebaseCredentialsFile)
	if err != nil {
		return Config{}, fmt.Errorf("resolve Firebase credentials path: %w", err)
	}
	if _, err := os.Stat(credentialsPath); err != nil {
		return Config{}, fmt.Errorf("Firebase credentials file not found at %q; download the service-account JSON and set FIREBASE_CREDENTIALS_FILE to its path", credentialsPath)
	}
	config.FirebaseCredentialsFile = credentialsPath

	return config, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
