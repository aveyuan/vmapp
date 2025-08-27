package user_repo

import (
	"context"
	"time"
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
	type Tab struct {
		ID   int
		Name string
	}
	var tb Tab
	if err := t.data.DB(ctx).WithContext(ctx).Table("uuc_sync_user").First(&tb).Error; err != nil {
		t.log.WithContext(ctx).Errorf("数据查询失败:%v", err)
		return ""
	}
	_, err := t.data.Redis.Set(ctx, "name", "zs", 100*time.Second).Result()
	if err != nil {
		t.log.WithContext(ctx).Errorf("redis 数据获取失败，%v", err)
	}
	t.log.WithContext(ctx).Info("redis 设置成功")
	// res, err := t.data.GreeterClient.SayHello(ctx, &v1.HelloRequest{Name: name})
	// if err != nil {
	// 	t.log.Errorf("调用服务失败:%v", err)
	// 	return ""
	// }
	return name
}

func (t *userRepo) GetHello(ctx context.Context) ([]models.User, error) {
	var list []models.User
	if err := t.data.DB(ctx).Model(new(models.User)).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
