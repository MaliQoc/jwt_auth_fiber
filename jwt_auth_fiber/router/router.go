package router

import (
	"github.com/gofiber/fiber/v2"
	"jwt_auth_fiber/handlers"
	"jwt_auth_fiber/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	tokenMiddleware := middleware.TokenMiddleware()

	v1.Post("/Login", handlers.Login)
	v1.Get("/protected", tokenMiddleware, handlers.Protected)
}

