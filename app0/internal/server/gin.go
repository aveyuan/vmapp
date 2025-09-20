package server

import (
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/service"

	"github.com/gin-gonic/gin"

)



func NewGin(c *conf.AppConf, bc *conf.BootComponent, userServer *service.SysService, data *base.Data) *gin.Engine {
	r := gin.New()
	r.POST("/login", userServer.Login)
	r.POST("/register", userServer.Register)
	r.POST("/forget", userServer.Forget)
	r.POST("/logout", userServer.LogOut)
	r.POST("/refertoken", userServer.ReferToken)
	return r
}
