package service

import (
	"context"
	"vmapp/api/app0/gf"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"

)

type GfUserService struct {
	c  *conf.AppConf
	uc *usecase.UserCase
	bc *conf.BootComponent
}

func NewGfUserService(c *conf.AppConf, bc *conf.BootComponent, uc *usecase.UserCase) *GfUserService {
	return &GfUserService{
		c:  c,
		uc: uc,
		bc: bc,
	}
}

func (t *GfUserService) Hello(ctx context.Context, req *gf.HelloReq) (res *gf.HelloRes, err error) {



	user, err := t.uc.GetUser(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	res = &gf.HelloRes{
		Content: user,
	}
	return
}
