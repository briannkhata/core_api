package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Environment string
	DatabaseURL string
	LogLevel    string
	JWTSecret   string
	Timeout     time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() *Config {
	// Load .env file
	godotenv.Load()

	return &Config{
		Port:        getEnv("PORT", "8081"),
		Environment: getEnv("ENVIRONMENT", "development"),
		DatabaseURL: getDatabaseURL(),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		JWTSecret:   getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
		Timeout:     getTimeout(),
	}
}

func getDatabaseURL() string {
	// If full DATABASE_URL is provided, use it
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return dbURL
	}

	// Otherwise, construct from individual components
	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "yathu_erp"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSLMode,
	)
}

func getTimeout() time.Duration {
	seconds := getEnvAsInt("TIMEOUT_SECONDS", 30)
	return time.Duration(seconds) * time.Second
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// IsDevelopment returns true if running in development mode
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsProduction returns true if running in production mode
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}
