package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Standard API response structure
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Pagination metadata
type PaginationMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// Paginated response
type PaginatedResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

// Logger interface for shared logging
type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
}

// SendSuccess sends a successful response
func SendSuccess(c *fiber.Ctx, message string, data interface{}) error {
	response := APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	return c.Status(http.StatusOK).JSON(response)
}

// SendError sends an error response
func SendError(c *fiber.Ctx, statusCode int, message string) error {
	response := APIResponse{
		Success: false,
		Message: message,
		Error:   message,
	}
	return c.Status(statusCode).JSON(response)
}

// SendPaginated sends a paginated response
func SendPaginated(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	response := PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	return c.Status(http.StatusOK).JSON(response)
}

// ValidateStruct validates a struct using validator
func ValidateStruct(s interface{}) error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		// Format validation errors
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, formatValidationError(err))
		}
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}

// formatValidationError formats a validation error message
func formatValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", err.Field(), err.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", err.Field(), err.Param())
	case "numeric":
		return fmt.Sprintf("%s must be numeric", err.Field())
	case "alpha":
		return fmt.Sprintf("%s must contain only alphabetic characters", err.Field())
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", err.Field())
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", err.Field())
	case "datetime":
		return fmt.Sprintf("%s must be a valid datetime", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}

// Helper functions for common operations

// ParseTime parses a time string with common formats
func ParseTime(timeStr string) (*time.Time, error) {
	formats := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.000Z",
		"2006-01-02 15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("invalid time format: %s", timeStr)
}

// FormatTime formats a time to ISO 8601 format
func FormatTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}

// StringPtr returns a pointer to a string
func StringPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to an int
func IntPtr(i int) *int {
	return &i
}

// Float64Ptr returns a pointer to a float64
func Float64Ptr(f float64) *float64 {
	return &f
}

// BoolPtr returns a pointer to a bool
func BoolPtr(b bool) *bool {
	return &b
}

// StringValue returns the value of a string pointer or empty string
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// IntValue returns the value of an int pointer or 0
func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// Float64Value returns the value of a float64 pointer or 0
func Float64Value(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// BoolValue returns the value of a bool pointer or false
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
