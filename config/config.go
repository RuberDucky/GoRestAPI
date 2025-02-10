package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type JWTConfig struct {
	SecretKey string
	ExpiresIn int // in hours
}

var AppConfig Config

func LoadConfig() {
	// Try to load .env file, but don't fail if it doesn't exist
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	AppConfig = Config{
		Server: ServerConfig{
			Port:    getEnvWithDefault("SERVER_PORT", "8080"),
			GinMode: getEnvWithDefault("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnvWithDefault("DB_HOST", "localhost"),
			User:     getEnvWithDefault("DB_USER", "postgres"),
			Password: getEnvWithDefault("DB_PASSWORD", "root"),
			DBName:   getEnvWithDefault("DB_NAME", "postgres"),
			Port:     getEnvWithDefault("DB_PORT", "5432"),
		},
		JWT: JWTConfig{
			SecretKey: getEnvWithDefault("JWT_SECRET_KEY", "your-secret-key"),
			ExpiresIn: 24, // 24 hours
		},
	}

	// Log the configuration being used
	log.Printf("Server Configuration: Port=%s, Mode=%s", AppConfig.Server.Port, AppConfig.Server.GinMode)
	log.Printf("Database Configuration: Host=%s, Port=%s, User=%s, DBName=%s", 
		AppConfig.Database.Host, 
		AppConfig.Database.Port, 
		AppConfig.Database.User, 
		AppConfig.Database.DBName)
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 