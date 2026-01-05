package middleware

import (
	"strings"
	"yathuerp/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse{
				Success: false,
				Message: "Authorization header required",
				Data:    nil,
			})
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}
			return []byte("your-secret-key"), nil // This should come from config
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse{
				Success: false,
				Message: "Invalid token",
				Data:    nil,
			})
		}

		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("userID", claims["user_id"])
			c.Locals("email", claims["email"])
		}

		return c.Next()
	}
}
