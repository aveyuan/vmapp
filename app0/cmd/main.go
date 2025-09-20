package main

import (
	"flag"

	"vmapp/app0/internal/conf"

	"github.com/aveyuan/vlogger"

	_ "go.uber.org/automaxprocs"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

const Version = "1.0.0"

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

// @Version 1.0.0
// @Title 后端服务
// @Description 提供API服务
// @ContactName aveyuan
// @ContactEmail aveyuan@163.com
// @ContactURL https://www.github.com/aveyuan
// @Server http://127.0.0.1 Server-1
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader 输入你的token
func main() {

	fc := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer fc.Close()

	if err := fc.Load(); err != nil {
		panic(err)
	}

	var ac conf.AppConf
	if err := fc.Scan(&ac); err != nil {
		panic(err)
	}

	if ac.Data != nil && ac.Data.DB != nil && ac.Data.DB.Logconfig != nil {
		ac.Data.DB.Logconfig.Level = ac.Logging.Level
	}

	ac.Logging.AppName = ac.App.AppName
	ac.Logging.AppVersion = Version
	ac.Logging.Env = ac.App.Env

	vlogger := vlogger.New(ac.Logging)

	bc := conf.NewBootComponent(&ac, vlogger)

	Handler, cleanup, err := wireGinApp(&ac, ac.Data, bc)
	if err != nil {
		panic(err)
	}

	Handler.Run(ac.Server.Http.Addr)

	defer cleanup()


	Handler.Run()

}
