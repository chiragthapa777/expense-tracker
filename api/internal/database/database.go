package database

import (
	"fmt"
	"time"

	"log"

	"github.com/chiragthapa777/expense-tracker-api/internal/config"
	customLogger "github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	myLogger := customLogger.GetLogger()
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // Use default Go logger
		logger.Config{
			SlowThreshold:             time.Second, // SQL query time threshold for "slow" queries
			LogLevel:                  logger.Info, // Log level: Silent, Error, Warn, Info
			IgnoreRecordNotFoundError: true,        // Ignore "record not found" errors
			Colorful:                  true,        // Enable colored logs
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		myLogger.Errorf("Failed to connect to database: %v", err)
		panic("Database connection failed")
	}

	DB = db
	myLogger.Info("Database connected successfully!")
}
