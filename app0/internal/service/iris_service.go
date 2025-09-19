package service

import (
	"context"
	"net/http"
	v1 "vmapp/api/app0/v1"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"
	"vmapp/pkg/vhttp"

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

// @Title 获取Hello
// @Param  .  query  v1.HelloReq  false  "请求参数"
// @Success  200  object  v1.HelloRes  "返回结果"
// @Tag users
// @Route /api/v1/hello [get]
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

// @Title 登录
// @Param  .  body  v1.LoginReq  true  "请求参数"
// @Success  200  object  v1.LoginRes  "返回结果"
// @Tag users
// @Route /api/v1/login [post]
func (t *IrisUserService) Login(c iris.Context) {

	var req v1.LoginReq
	if err := c.ReadJSON(&req); err != nil {
		vhttp.ErrorHandle(c, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	resd, err := t.uc.Login(c, &req)
	if err != nil {
		vhttp.ErrorHandle(c, err)
		return
	}

	c.JSON(resd)

}
