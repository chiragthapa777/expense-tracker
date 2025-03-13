package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

// LogLevel represents the severity of the log
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

// Logger is our custom logger structure
type Logger struct {
	logger *log.Logger
}

// singleton instance and sync.Once for thread-safe initialization
var (
	loggerInstance *Logger
	once           sync.Once
)

// GetLogger returns the singleton logger instance
func GetLogger() *Logger {
	once.Do(func() {

		// Ensure logs directory exists
		err := os.MkdirAll("./logs", 0755)
		if err != nil {
			log.Fatalf("Failed to create logs directory: %v", err)
		}

		// Create log file
		file, err := os.OpenFile("./logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		// Log to both file and console
		mw := io.MultiWriter(os.Stdout, file)
		loggerInstance = &Logger{
			logger: log.New(mw, "", 0),
		}
	})
	return loggerInstance
}

// formatMessage adds timestamp and level to the log message
func (l *Logger) formatMessage(level LogLevel, message string) string {
	now := time.Now().Format("2006-01-02 15:04:05")
	var levelStr string
	switch level {
	case INFO:
		levelStr = "INFO"
	case WARN:
		levelStr = "WARN"
	case ERROR:
		levelStr = "ERROR"
	}
	return fmt.Sprintf("[%s] [%s] %s", now, levelStr, message)
}

// Info logs an informational message
func (l *Logger) Info(message string) {
	l.logger.Println(l.formatMessage(INFO, message))
}

// Warn logs a warning message
func (l *Logger) Warn(message string) {
	l.logger.Println(l.formatMessage(WARN, message))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Println(l.formatMessage(WARN, fmt.Sprintf(format, args...)))
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.logger.Println(l.formatMessage(ERROR, message))
}

// Errorf logs an error with formatting
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Println(l.formatMessage(ERROR, fmt.Sprintf(format, args...)))
}
