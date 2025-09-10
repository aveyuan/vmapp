package usecase

import (
	"context"
	"errors"
	v1 "vmapp/api/app0/v1"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/models"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/os/gctx"
)

type UserCase struct {
	up  repo.UserRepo
	tx  repo.Transaction
	log *log.Helper
}

func NewUserCase(up repo.UserRepo, tx repo.Transaction, bc *conf.BootComponent) *UserCase {
	return &UserCase{
		up:  up,
		tx:  tx,
		log: bc.Logger,
	}
}

func (t *UserCase) CreateUser(ctx context.Context) {

}

func (t *UserCase) ListUser(ctx context.Context) {

}

func (t *UserCase) UpdateUser(ctx context.Context) {

}

func (t *UserCase) DelectUser(ctx context.Context) {

}

func (t *UserCase) GetUser(ctx context.Context, name string) (string, error) {

	res := t.up.GetUser(ctx, name)
	if res == "张三" {
		t.log.WithContext(ctx).Error("名字错误")
		return "", errors.New("姓名错误")
	}
	return res, nil
}

// GetHello 测试
func (t *UserCase) GetHello(ctx context.Context) ([]models.User, error) {
	t.log.Info(gctx.CtxId(ctx))
	m, err := t.up.GetHello(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取数据失败,%+v", err)
		return nil, errors.New("数据读取失败")
	}
	return m, nil
}

func (t *UserCase) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {

	res = &v1.LoginRes{
		Token:  "123",
		Expire: 1000,
	}
	return
}
