package server

import (
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Response struct {
	Message string      `json:"message" dc:"消息提示"`
	Data    interface{} `json:"data"    dc:"执行结果"`
}

func ResponseMiddleware(r *ghttp.Request) {
	r.SetCtx(gi18n.WithLanguage(r.GetCtx(), "zh-CN"))

	r.Middleware.Next()

	var (
		msg string
		res = r.GetHandlerResponse()
		err = r.GetError()
	)
	if err != nil {
		msg = err.Error()
	} else {
		msg = "OK"
	}
	r.Response.WriteJson(Response{
		Message: msg,
		Data:    res,
	})
}

func NewGf(c *conf.AppConf, bc *conf.BootComponent, userServer *service.GfUserService, data *base.Data) *ghttp.Server {
	r := g.Server()

	r.Group("/hello", func(group *ghttp.RouterGroup) {
		group.Middleware(ResponseMiddleware)
		group.Bind(
			userServer,
		)
	})
	r.SetOpenApiPath("/api.json")
	r.SetSwaggerPath("/swagger")
	r.SetPort(8000)
	return r
}
