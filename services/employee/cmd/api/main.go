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

	"yathu_erp/services/employee-service/internal/infrastructure/http"
	"yathu_erp/services/employee-service/internal/infrastructure/persistence/postgres"

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
	if err := database.RunMigrations(db, "employee-service"); err != nil {
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
			"service": "employee-service",
			"version": "1.0.0",
		})
	})

	// Setup routes
	employeeRepo := postgres.NewRepository(db, logger)
	http.SetupRoutes(app, employeeRepo)

	// Graceful shutdown
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		logger.Info("Shutting down employee-service...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(ctx); err != nil {
			logger.Error("Error during shutdown:", err)
		}
	}()

	// Start server
	logger.Info("Employee service starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
