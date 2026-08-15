package service

import (
	"context"
	"time"

	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/models"
	"github.com/Harsh-karn/Go-Backend-Development-Task/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error)
	GetUser(ctx context.Context, id int32) (models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context, limit, offset int32) ([]models.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func CalculateAge(dob, now time.Time) int {
	age := now.Year() - dob.Year()
	// If birthday hasn't occurred this year yet, subtract 1
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	return age
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.CreateUser(ctx, req.Name, dob)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *userService) GetUser(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	age := CalculateAge(user.Dob, time.Now())
	
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  &age,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.UpdateUser(ctx, id, req.Name, dob)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context, limit, offset int32) ([]models.UserResponse, error) {
	users, err := s.repo.ListUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []models.UserResponse
	now := time.Now()
	for _, u := range users {
		age := CalculateAge(u.Dob, now)
		responses = append(responses, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Format("2006-01-02"),
			Age:  &age,
		})
	}

	// Return empty array instead of null when no users found
	if responses == nil {
		responses = []models.UserResponse{}
	}

	return responses, nil
}
