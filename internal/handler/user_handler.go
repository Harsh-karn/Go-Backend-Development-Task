package handler

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/logger"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/models"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	svc      service.UserService
	validate *validator.Validate
}

func NewUserHandler(svc service.UserService) *UserHandler {
	v := validator.New()
	
	// Custom validation for date logic
	_ = v.RegisterValidation("datetime", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		format := fl.Param()
		parsedDate, err := time.Parse(format, dateStr)
		if err != nil {
			return false
		}
		// Ensure date is not in the future
		if parsedDate.After(time.Now()) {
			return false
		}
		return true
	})

	return &UserHandler{
		svc:      svc,
		validate: v,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Warn("Invalid request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Invalid JSON format"})
	}

	if err := h.validate.Struct(req); err != nil {
		logger.Log.Warn("Validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Validation failed: " + err.Error()})
	}

	res, err := h.svc.CreateUser(c.Context(), req)
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to create user"})
	}

	logger.Log.Info("User created", zap.Int32("id", res.ID))
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Invalid user ID"})
	}

	res, err := h.svc.GetUser(c.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "User not found"})
		}
		logger.Log.Error("Failed to get user", zap.Error(err), zap.Int("id", id))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to get user"})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Invalid user ID"})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Invalid JSON format"})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Validation failed: " + err.Error()})
	}

	// Check if user exists first
	_, err = h.svc.GetUser(c.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to verify user"})
	}

	res, err := h.svc.UpdateUser(c.Context(), int32(id), req)
	if err != nil {
		logger.Log.Error("Failed to update user", zap.Error(err), zap.Int("id", id))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to update user"})
	}

	logger.Log.Info("User updated", zap.Int32("id", res.ID))
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Invalid user ID"})
	}

	// Check if user exists first
	_, err = h.svc.GetUser(c.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to verify user"})
	}

	if err := h.svc.DeleteUser(c.Context(), int32(id)); err != nil {
		logger.Log.Error("Failed to delete user", zap.Error(err), zap.Int("id", id))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to delete user"})
	}

	logger.Log.Info("User deleted", zap.Int("id", id))
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	res, err := h.svc.ListUsers(c.Context(), int32(limit), int32(offset))
	if err != nil {
		logger.Log.Error("Failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to list users"})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
