package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"golang-api/config"
)

type errorResponse struct {
	Error string `json:"error"`
}

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.SECRET),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func IsAdmin(context *fiber.Ctx) error {
	user := context.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	if role == "admin" {
		return context.Next()
	}
	return context.Status(403).JSON(errorResponse{Error: "You don't have access to this route"})
}