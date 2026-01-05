package routes

import (
	"yathuerp/handlers/employees"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "yathuerp-api",
		})
	})

	// API versioning
	api := app.Group("/api/v1")

	// Employees module routes
	employeeHandler := employees.NewHandler(db)
	employeesGroup := api.Group("/employees")
	{
		employeesGroup.Get("/", employeeHandler.GetAllEmployees)
		employeesGroup.Get("/:id", employeeHandler.GetEmployeeByID)
		employeesGroup.Post("/", employeeHandler.CreateEmployee)
		employeesGroup.Put("/:id", employeeHandler.UpdateEmployee)
		employeesGroup.Delete("/:id", employeeHandler.DeleteEmployee)
	}
}
