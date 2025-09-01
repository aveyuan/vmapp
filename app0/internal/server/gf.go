package server

import (
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/service"
	"vmapp/pkg/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
)

func I18nMiddleware(r *ghttp.Request) {
	r.SetCtx(gi18n.WithLanguage(r.GetCtx(), "zh-CN"))

}

func NewGf(c *conf.AppConf, bc *conf.BootComponent, userServer *service.GfUserService, data *base.Data) *ghttp.Server {
	g.SetDebug(false)
	r := g.Server()
	r.SetGraceful(true)
	// r.SetLogger(&glog.Logger{})
	// r.SetSessionMaxAge(time.Minute)
	// r.SetSessionStorage(gsession.NewStorageMemory())
	// r.Use(I18nMiddleware)
	

	// r.Use(ghttp.MiddlewareGzip)
	// r.Use(middleware.NewRequestLog(bc.Logger))
	// r.Use(middleware.NewRecoverLog(bc.Logger))
	r.Use(middleware.MiddlewareHandlerResponse)

	r.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(
			userServer,
		)
	})
	// r.SetOpenApiPath("/api.json")
	// r.SetSwaggerPath("/swagger")
	r.SetPort(8000)
	return r
}
