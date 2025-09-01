package service

import (
	"context"
	v1 "vmapp/api/app0/v1"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"

	"github.com/gin-gonic/gin"
)

type GinUserService struct {
	c  *conf.AppConf
	uc *usecase.UserCase
	bc *conf.BootComponent
}

func NewGinUserService(c *conf.AppConf, bc *conf.BootComponent, uc *usecase.UserCase) *GinUserService {
	return &GinUserService{
		c:  c,
		uc: uc,
		bc: bc,
	}
}

func (t *GinUserService) Hello(c *gin.Context) {

	var req v1.HelloReq
	var res v1.HelloRes
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resd, err := t.uc.GetUser(context.Background(), req.Name)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res.Content = resd
	c.JSON(200, res)


}
