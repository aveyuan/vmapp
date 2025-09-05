package server

import (
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/service"

	"github.com/kataras/iris/v12"
)

func NewIris(c *conf.AppConf, bc *conf.BootComponent, userServer *service.IrisUserService, data *base.Data) *iris.Application {
	r := iris.New()

	api := r.Party("/api")
	v1 := api.Party("/v1")
	v1.Get("/hello", userServer.Hello)

	return r
}
