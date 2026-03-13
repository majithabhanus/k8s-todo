package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	DB_HOST    string
	DB_PORT    string
)

func LoadConfig() {
	// Load .env ONLY if present (local dev)
	_ = godotenv.Load() // safe for local, ignored in Docker

	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")
	JWTSecret = os.Getenv("JWT_SECRET")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
}
