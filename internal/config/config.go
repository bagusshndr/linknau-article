package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
}

func Load() *Config {
	_ = godotenv.Load() // ignore error kalau .env ga ada

	cfg := &Config{
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASSWORD", ""),
		DBName:    getEnv("DB_NAME", "linknau_articles"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}

	if cfg.DBUser == "" {
		log.Println("WARNING: DB_USER is empty")
	}

	return cfg
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}
