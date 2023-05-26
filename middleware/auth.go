package middleware

import (
	"github.com/MrRytis/go-fiber-test/service"
	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-jwt/jwt/v5"
	"strings"
)

func (middleware Middleware) Auth(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	if tokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
	}

	claims, err := service.ParseJWT(tokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// Check if the token is blacklisted (User logged out)
	if service.IsBlacklisted(middleware.cache, tokenString) {
		return fiber.NewError(fiber.StatusUnauthorized, "Token is blacklisted")
	}

	c.Locals("uid", claims["uid"].(string))
	c.Locals("expiresAt", int64(claims["exp"].(float64)))
	c.Locals("jwt", tokenString)

	return c.Next()
}
