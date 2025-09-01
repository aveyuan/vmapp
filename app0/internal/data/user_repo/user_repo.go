package user_repo

import (
	"context"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/models"

	"vmapp/app0/internal/data/base"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *base.Data
	log  *log.Helper
}

func NewUserRepo(data *base.Data, component *conf.BootComponent) repo.UserRepo {
	return &userRepo{
		data: data,
		log:  component.Logger,
	}
}

func (t *userRepo) CreateUser(ctx context.Context, name string) {
	log.Info("hello", name)
}

func (t *userRepo) UpdateUser(ctx context.Context, name string) {
	log.Info("hello", name)
}

func (t *userRepo) DeleteUser(ctx context.Context, name string) {
	log.Info("hello", name)
}

func (t *userRepo) ListUser(ctx context.Context, name string) {
	log.Info("hello", name)
}

func (t *userRepo) GetUser(ctx context.Context, name string) string {
	return name
}

func (t *userRepo) GetHello(ctx context.Context) ([]models.User, error) {
	var list []models.User
	if err := t.data.DB(ctx).Model(new(models.User)).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
