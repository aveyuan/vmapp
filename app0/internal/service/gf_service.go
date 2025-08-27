package service

import (
	"context"
	v1 "vmapp/api/app0/v1"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"

	"github.com/gogf/gf/v2/frame/g"
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

func (t *GfUserService) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {

	if err := g.Validator().Data(req).I18n(g.I18n("zh-CN")).Run(ctx); err != nil {

		g.Dump(err.String())
	}

	res = &v1.HelloRes{
		Content: "hello world",
	}
	return
}
