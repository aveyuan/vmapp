package server

import (
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/service"

	"github.com/gin-gonic/gin"

)



func NewGin(c *conf.AppConf, bc *conf.BootComponent, userServer *service.GinUserService, data *base.Data) *gin.Engine {
	r := gin.New()
	r.GET("/hello", userServer.Hello)

	return r
}
