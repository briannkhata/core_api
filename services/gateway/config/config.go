package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	Services map[string]string
}

func Load() *Config {
	// Load .env file
	godotenv.Load()

	return &Config{
		Port: getEnv("GATEWAY_PORT", "8080"),
		Services: map[string]string{
			"employees":   getEnv("EMPLOYEES_SERVICE_URL", "http://localhost:8081"),
			"payroll":     getEnv("PAYROLL_SERVICE_URL", "http://localhost:8082"),
			"attendance":  getEnv("ATTENDANCE_SERVICE_URL", "http://localhost:8083"),
			"leave":       getEnv("LEAVE_SERVICE_URL", "http://localhost:8084"),
			"loans":       getEnv("LOANS_SERVICE_URL", "http://localhost:8085"),
			"users":       getEnv("USERS_SERVICE_URL", "http://localhost:8086"),
			"performance": getEnv("PERFORMANCE_SERVICE_URL", "http://localhost:8087"),
			"banking":     getEnv("BANKING_SERVICE_URL", "http://localhost:8088"),
			"audit":       getEnv("AUDIT_SERVICE_URL", "http://localhost:8089"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
