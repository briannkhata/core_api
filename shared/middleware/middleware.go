package middleware

import (
	"time"

	"yathuerp/shared/logger"
	"yathuerp/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestID middleware adds a unique request ID to each request
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("X-Request-ID", requestID)
		c.Locals("request_id", requestID)
		return c.Next()
	}
}

// Logger middleware logs request details
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Log request details
		logger.Info("Request processed",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"duration", duration,
			"request_id", c.Locals("request_id"),
			"user_agent", c.Get("User-Agent"),
			"ip", c.IP(),
		)

		return err
	}
}

// Tracing middleware adds distributed tracing headers
func Tracing() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Add trace headers
		traceID := c.Get("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}

		spanID := uuid.New().String()

		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)
		c.Locals("trace_id", traceID)
		c.Locals("span_id", spanID)

		return c.Next()
	}
}

// ErrorHandler provides centralized error handling
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Log the error
	logger.Error("Request error",
		"error", err,
		"method", c.Method(),
		"path", c.Path(),
		"request_id", c.Locals("request_id"),
	)

	// Handle different types of errors
	if e, ok := err.(*fiber.Error); ok {
		return utils.SendError(c, e.Code, e.Message)
	}

	// Default error response
	return utils.SendError(c, fiber.StatusInternalServerError, "Internal server error")
}

// RequestBodyLogger middleware logs request body for debugging
func RequestBodyLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Only log for POST/PUT/PATCH requests
		if c.Method() != "POST" && c.Method() != "PUT" && c.Method() != "PATCH" {
			return c.Next()
		}

		// Read body
		body := c.Body()

		// Log body (be careful with sensitive data)
		logger.Debug("Request body",
			"method", c.Method(),
			"path", c.Path(),
			"body", string(body),
			"request_id", c.Locals("request_id"),
		)

		// Restore body for next handlers
		c.Request().SetBody(body)

		return c.Next()
	}
}

// ResponseLogger middleware logs response details
func ResponseLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Capture response
		err := c.Next()

		// Log response
		logger.Debug("Response",
			"status", c.Response().StatusCode(),
			"content_length", len(c.Response().Body()),
			"request_id", c.Locals("request_id"),
		)

		return err
	}
}

// RateLimiter middleware (basic implementation)
func RateLimiter() fiber.Handler {
	// TODO: Implement proper rate limiting with Redis or in-memory store
	return func(c *fiber.Ctx) error {
		// For now, just pass through
		return c.Next()
	}
}
