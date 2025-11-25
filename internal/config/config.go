package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	DBSSL   string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASSWORD", ""),
		DBName:  getEnv("DB_NAME", "linknau_articles"),
		DBSSL:   getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
