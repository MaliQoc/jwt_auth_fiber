package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"jwt_auth_fiber/models"
	"jwt_auth_fiber/utils"
)

func TokenMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(models.Response{Message: "unauthorized", Status: http.StatusUnauthorized})
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(models.Response{Message: err.Error(), Status: http.StatusUnauthorized})
		}

		return c.Next()
	}
}

