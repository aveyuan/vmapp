package repo

import (
	"context"
	"vmapp/app0/internal/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, name string)
	ListUser(ctx context.Context, name string)
	UpdateUser(ctx context.Context, name string)
	DeleteUser(ctx context.Context, name string)
	GetUser(ctx context.Context, name string) string
	GetHello(ctx context.Context) ([]models.User, error)
}
