package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	Environment string
	LogLevel    string
	Server      ServerConfig
	Agent       AgentConfig
	Database    DatabaseConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Host string
	Port int
}

// AgentConfig holds agent-specific configuration
type AgentConfig struct {
	MaxConcurrent int
	Timeout       int // seconds
	ModelProvider string
	APIKey        string
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Type     string
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvInt("SERVER_PORT", 8080),
		},
		Agent: AgentConfig{
			MaxConcurrent: getEnvInt("AGENT_MAX_CONCURRENT", 10),
			Timeout:       getEnvInt("AGENT_TIMEOUT", 300),
			ModelProvider: getEnv("MODEL_PROVIDER", "openai"),
			APIKey:        getEnv("MODEL_API_KEY", ""),
		},
		Database: DatabaseConfig{
			Type:     getEnv("DB_TYPE", "postgres"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			Name:     getEnv("DB_NAME", "ade"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}

	if c.Agent.MaxConcurrent < 1 {
		return fmt.Errorf("invalid agent max concurrent: %d", c.Agent.MaxConcurrent)
	}

	if c.Agent.Timeout < 1 {
		return fmt.Errorf("invalid agent timeout: %d", c.Agent.Timeout)
	}

	return nil
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
