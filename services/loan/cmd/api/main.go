package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"yathu_erp/shared/config"
	"yathu_erp/shared/database"
	"yathu_erp/shared/logger"
	"yathu_erp/shared/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	logger.Init(cfg.Environment)

	// Initialize database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.RunMigrations(db, "loan-service"); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())
	app.Use(middleware.Tracing())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "loan-service",
			"version": "1.0.0",
		})
	})

	// Setup routes
	setupRoutes(app, db)

	// Graceful shutdown
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		logger.Info("Shutting down loan-service...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(ctx); err != nil {
			logger.Error("Error during shutdown:", err)
		}
	}()

	// Start server
	logger.Info("Loan service starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}

func setupRoutes(app *fiber.App, db interface{}) {
	// TODO: Setup routes from internal/infrastructure/http
	// This will be implemented in next step
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":    "YathuERP Loan Service is running",
			"service":    "loan-service",
			"go_version": "1.22",
		})
	})
}
