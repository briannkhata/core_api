package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Init initializes the global logger based on environment
func Init(environment string) {
	var config zap.Config

	switch environment {
	case "production":
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "development":
		config = zap.NewDevelopmentConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	default:
		config = zap.NewDevelopmentConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	// Customize encoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	Logger, err = config.Build()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Replace global logger
	zap.ReplaceGlobals(Logger)
}

// Info logs an info message
func Info(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Sugar().Infow(msg, args...)
	} else {
		log.Printf("INFO: %s %v", msg, args)
	}
}

// Error logs an error message
func Error(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Sugar().Errorw(msg, args...)
	} else {
		log.Printf("ERROR: %s %v", msg, args)
	}
}

// Debug logs a debug message
func Debug(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Sugar().Debugw(msg, args...)
	} else {
		log.Printf("DEBUG: %s %v", msg, args)
	}
}

// Warn logs a warning message
func Warn(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Sugar().Warnw(msg, args...)
	} else {
		log.Printf("WARN: %s %v", msg, args)
	}
}

// Sync flushes any buffered log entries
func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
}

// Close closes the logger
func Close() {
	if Logger != nil {
		Logger.Sync()
	}
}
