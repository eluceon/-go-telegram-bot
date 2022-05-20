package repository

import (
	"context"
	"errors"
	"github.com/eluceon/quizwithmebot/internal/app/server/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = errors.New("not found")

type repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}

type Repository interface {
	CreateUser(context.Context, models.User) error
	ReadUser(context.Context, int64) (models.User, error)
	UpdateUser(context.Context, models.User) error
	DeleteUser(context.Context, int64) error
	GetTopUsers(context.Context) ([]models.User, error)
}
