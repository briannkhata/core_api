package middleware

import (
	"log"
	"yathuerp/utils"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Default status code
	code := fiber.StatusInternalServerError

	// Check if it's a fiber error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Log the error
	log.Printf("Error: %s", err.Error())

	// Return JSON response
	return c.Status(code).JSON(utils.APIResponse{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
}
