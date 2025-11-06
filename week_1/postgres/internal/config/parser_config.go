package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	GRPCHost   string
	GRPCPort   string
	Env        string
}

func LoadConfig(env string, service string) *Config {
	if service == "" {
		service = "auth"
	}

	if env == "" {
		env = "dev"
	}

	envFile := "/" + service + "/" + service + "_config_" + env + ".env"

	if err := godotenv.Load("internal/config/" + envFile); err != nil {
		log.Printf("⚠️  Warning: could not load %s, using defaults", envFile)
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "auth"),
		DBUser:     getEnv("DB_USER", "auth"),
		DBPassword: getEnv("DB_PASSWORD", "qwerty1"),
		GRPCHost:   getEnv("GRPC_HOST", "localhost"),
		GRPCPort:   getEnv("GRPC_PORT", "50051"),
		Env:        getEnv("ENV", env),
	}
}

func getEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}
