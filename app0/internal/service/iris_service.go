package service

import (
	"context"
	v1 "vmapp/api/app0/v1"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"

	"github.com/kataras/iris/v12"
)

type IrisUserService struct {
	c  *conf.AppConf
	uc *usecase.UserCase
	bc *conf.BootComponent
}

func NewIrisUserService(c *conf.AppConf, bc *conf.BootComponent, uc *usecase.UserCase) *IrisUserService {
	return &IrisUserService{
		c:  c,
		uc: uc,
		bc: bc,
	}
}

func (t *IrisUserService) Hello(c iris.Context) {

	var req v1.HelloReq
	var res v1.HelloRes
	if err := c.ReadQuery(&req); err != nil {
		c.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}

	resd, err := t.uc.GetUser(context.Background(), req.Name)
	if err != nil {
		c.JSON(iris.Map{
			"error": err.Error(),
		})
	}

	res.Content = resd
	c.JSON(res)

}
