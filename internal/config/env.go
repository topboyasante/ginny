package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var ENV = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		ServerPort: getEnv("SERVER_PORT", "4000"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBName:     getEnv("DB_NAME", "ticketly"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
