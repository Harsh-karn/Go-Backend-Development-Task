package main

import (
	"database/sql"
	"log"

	"github.com/Harsh-karn/Go-Backend-Development-Task/config"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/handler"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/logger"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/repository"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/routes"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/service"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	defer logger.Sync()

	logger.Log.Info("Starting server...")

	// Load Config
	cfg := config.LoadConfig()

	// Connect to Database
	db, err := sql.Open("postgres", cfg.DBDSN)
	if err != nil {
		logger.Log.Fatal("Cannot connect to database", zap.Error(err))
	}
	if err = db.Ping(); err != nil {
		logger.Log.Fatal("Cannot ping database", zap.Error(err))
	}
	logger.Log.Info("Connected to PostgreSQL database successfully")

	// Initialize Repository, Service, and Handler
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Initialize Fiber App
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Log.Error("Fiber error", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Setup Routes
	routes.SetupRoutes(app, userHandler)

	// Start Server
	logger.Log.Info("Server listening on port " + cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		logger.Log.Fatal("Server failed", zap.Error(err))
	}
}
