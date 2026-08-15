package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Harsh-karn/Go-Backend-Development-Task/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string, dob time.Time) (db.User, error)
	GetUser(ctx context.Context, id int32) (db.User, error)
	UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (db.User, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error)
}

type userRepository struct {
	querier db.Querier
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{
		querier: db.New(conn),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, name string, dob time.Time) (db.User, error) {
	return r.querier.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *userRepository) GetUser(ctx context.Context, id int32) (db.User, error) {
	return r.querier.GetUser(ctx, id)
}

func (r *userRepository) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (db.User, error) {
	return r.querier.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.querier.DeleteUser(ctx, id)
}

func (r *userRepository) ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error) {
	return r.querier.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}
