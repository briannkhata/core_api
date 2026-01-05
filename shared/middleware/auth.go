package middleware

import (
	"fmt"
	"strings"
	"time"

	"yathuerp/shared/logger"
	"yathuerp/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string   `json:"user_id"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

type AuthMiddleware struct {
	logger logger.Logger
}

func NewAuthMiddleware(logger logger.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		logger: logger,
	}
}

func (am *AuthMiddleware) JWTAuth(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip auth for health check and login endpoints
		if c.Path() == "/health" || c.Path() == "/api/v1/auth/login" {
			return c.Next()
		}

		// Get token from header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SendError(c, fiber.StatusUnauthorized, "Missing authorization header")
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return utils.SendError(c, fiber.StatusUnauthorized, "Invalid authorization header format")
		}

		// Parse and validate token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			am.logger.Error("Invalid token", "error", err)
			return utils.SendError(c, fiber.StatusUnauthorized, "Invalid token")
		}

		// Extract claims
		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			return utils.SendError(c, fiber.StatusUnauthorized, "Invalid token claims")
		}

		// Store user info in context
		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)
		c.Locals("roles", claims.Roles)

		am.logger.Info("User authenticated", "user_id", claims.UserID, "email", claims.Email)
		return c.Next()
	}
}

// RoleBasedAuth middleware for role-based access control
func (am *AuthMiddleware) RoleBasedAuth(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRoles, exists := c.Locals("roles").([]string)
		if !exists {
			return utils.SendError(c, fiber.StatusForbidden, "No roles found in context")
		}

		// Check if user has any of the allowed roles
		for _, allowedRole := range allowedRoles {
			for _, userRole := range userRoles {
				if userRole == allowedRole {
					return c.Next()
				}
			}
		}

		am.logger.Error("Access denied - insufficient roles", "user_roles", userRoles, "required_roles", allowedRoles)
		return utils.SendError(c, fiber.StatusForbidden, "Insufficient permissions")
	}
}

// GenerateJWT creates a new JWT token
func GenerateJWT(userID, email string, roles []string, secret string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			// TODO: Set proper expiration
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
