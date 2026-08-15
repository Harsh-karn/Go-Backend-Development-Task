package routes

import (
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/handler"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	// Apply global middleware
	app.Use(middleware.RequestLogger())

	// API Group
	api := app.Group("/users")

	api.Post("/", userHandler.CreateUser)
	api.Get("/", userHandler.ListUsers)
	api.Get("/:id", userHandler.GetUser)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)
}
