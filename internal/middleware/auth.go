package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	//"os"
	"speakbuddy/pkg/utils"
)

func AuthRequired(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "missing token"})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{"error": "invalid auth format"})
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return utils.JwtSecret(), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}
