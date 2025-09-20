package server

import (
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/service"

	"github.com/gin-gonic/gin"
)

func NewGin(c *conf.AppConf, bc *conf.BootComponent, userServer *service.SysService, data *base.Data) *gin.Engine {
	r := gin.New()
	// 公共API
	api := r.Group("/api")
	api.POST("/login", userServer.Login)
	api.POST("/register", userServer.Register)
	api.POST("/forget", userServer.Forget)
	api.POST("/getcode", userServer.GetCode)
	api.POST("/captcha", userServer.Captcha)
	// 系统API
	sys := api.Group("/sys")
	sys.POST("/logout", userServer.LogOut)
	sys.POST("/refertoken", userServer.ReferToken)
	sys.POST("/restart", userServer.Restart)
	sys.POST("/update_userpass", userServer.UpdateUserPass)
	sys.POST("/update_userprofile", userServer.UpdateUserProfile)

	return r
}
