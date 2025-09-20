//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"vmapp/app0/internal/biz"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data"
	"vmapp/app0/internal/server"
	"vmapp/app0/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)


func wireGinApp(ac *conf.AppConf, dc *conf.Data, bc *conf.BootComponent) (*gin.Engine, func(), error) {
	panic(wire.Build(service.ProviderService, server.ProviderServer, biz.ProviderBiz, data.ProviderData))
}
