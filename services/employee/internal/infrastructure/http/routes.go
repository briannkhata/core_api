package http

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db interface{}) {
	// API versioning
	api := app.Group("/api/v1")

	// Employee CRUD routes
	employees := api.Group("/employees")
	{
		employees.Get("/", handlers.GetAllEmployees)
		employees.Get("/:id", handlers.GetEmployeeByID)
		employees.Post("/", handlers.CreateEmployee)
		employees.Put("/:id", handlers.UpdateEmployee)
		employees.Delete("/:id", handlers.DeleteEmployee)
	}

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "employee-service",
			"version": "1.0.0",
		})
	})
}
