package main

import (
	"flag"

	"vmapp/app0/internal/conf"

	"github.com/aveyuan/vlogger"
	"github.com/kataras/iris/v12"

	_ "go.uber.org/automaxprocs"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

const Version = "1.0.0"

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

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

	vlogger := vlogger.New(ac.Logging)

	bc := conf.NewBootComponent(vlogger)

	Handler, cleanup, err := wireIrisApp(&ac, bc)
	if err != nil {
		panic(err)
	}

	defer cleanup()

	Handler.Run(iris.Addr(ac.Server.Http.Addr))


}
