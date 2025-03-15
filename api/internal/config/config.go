package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Port                 string `validate:"required,numeric"` // Server port (required, but has default)
	DBHost               string `validate:"required"`         // Database host (required, has default)
	DBPort               string `validate:"required,numeric"` // Database port (required, has default)
	DBUser               string `validate:"required"`         // Database user (required, no default)
	DBPass               string `validate:"required"`         // Database password (required, no default)
	DBName               string `validate:"required"`         // Database name (required, no default)
	JWTSecret            string `validate:"required"`         // Database name (required, no default)
	S3_API               string `validate:"required"`
	R2_TOKEN             string `validate:"required"`
	R2_ACCESS_KEY_ID     string `validate:"required"`
	R2_SECRET_ACCESS_KEY string `validate:"required"`
	R2_BUCKET_NAME       string `validate:"required"`
	R2_ACCOUNT_ID        string `validate:"required"`
}

var (
	config Config
	once   sync.Once
)

func GetConfig() Config {
	once.Do(func() {
		config = loadConfig()
	})
	return config
}

// LoadConfig loads and validates the configuration
func loadConfig() Config {
	log := logger.GetLogger()
	// Try to load .env file
	err := godotenv.Load()
	if err != nil {
		log.Error("No .env file found - this is required!")
		panic("Configuration failed: .env file is missing")
	}

	// Initialize config with defaults for optional fields
	cfg := Config{
		Port:   "3000",      // Default for Port
		DBHost: "localhost", // Default for DBHost
		DBPort: "5432",      // Default for DBPort
	}

	// Load environment variables into config
	cfg.Port = getEnv("PORT", cfg.Port)
	cfg.DBHost = getEnv("DB_HOST", cfg.DBHost)
	cfg.DBPort = getEnv("DB_PORT", cfg.DBPort)
	cfg.DBUser = getEnv("DB_USER", "")
	cfg.DBPass = getEnv("DB_PASS", "")
	cfg.DBName = getEnv("DB_NAME", "")
	cfg.JWTSecret = getEnv("JWT_SECRET", "")
	cfg.S3_API = getEnv("S3_API", "")
	cfg.R2_TOKEN = getEnv("R2_TOKEN", "")
	cfg.R2_ACCESS_KEY_ID = getEnv("R2_ACCESS_KEY_ID", "")
	cfg.R2_SECRET_ACCESS_KEY = getEnv("R2_SECRET_ACCESS_KEY", "")
	cfg.R2_BUCKET_NAME = getEnv("R2_BUCKET_NAME", "")
	cfg.R2_ACCOUNT_ID = getEnv("R2_ACCOUNT_ID", "")

	// Validate the config using the validator
	if err := validateConfig(cfg, log); err != nil {
		log.Errorf("Configuration validation failed: %v", err)
		panic(fmt.Sprintf("Configuration error: %v", err))
	}

	log.Info("Configuration loaded successfully")
	return cfg
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}

// validateConfig uses the validator library to check the config
func validateConfig(cfg Config, log *logger.Logger) error {
	validate := validator.New()
	err := validate.Struct(cfg)
	if err != nil {
		// If validation fails, build a detailed error message
		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			errMsg += fmt.Sprintf("Field '%s' failed validation: %s; ", err.Field(), err.Tag())
		}
		return fmt.Errorf(errMsg)
	}
	return nil
}
