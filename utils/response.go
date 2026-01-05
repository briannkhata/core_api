package utils

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func SendSuccess(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
	})
}

func SendPaginated(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.JSON(PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}
