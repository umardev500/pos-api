package pkg

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearerToken := c.Get("Authorization")
		if bearerToken == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Chek for the Bearer prefix
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(bearerToken, bearerPrefix) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Extract the token
		tokenString := strings.TrimPrefix(bearerToken, bearerPrefix)
		secret := os.Getenv("JWT_SECRET")
		claims, err := ValidateJWT(tokenString, secret)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		c.Locals("user_id", (*claims)["user_id"])

		return c.Next()
	}
}
