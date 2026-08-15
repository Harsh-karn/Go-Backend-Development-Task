package middleware

import (
	"time"

	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// RequestLogger middleware injects request ID and logs duration
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		reqID := c.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		c.Set("X-Request-ID", reqID)
		
		// Add request ID to locals so it could potentially be extracted by handlers if needed
		c.Locals("requestid", reqID)

		err := c.Next()

		duration := time.Since(start)

		logger.Log.Info("HTTP Request",
			zap.String("request_id", reqID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)

		return err
	}
}
