package conf

import (
	"github.com/aveyuan/vbasedata"

	"github.com/aveyuan/vlogger"
)

type AppConf struct {
	App     *vbasedata.App     `yaml:"app" json:"app"`
	Server  *Server            `yaml:"server" json:"server"`
	Data    *Data              `yaml:"data" json:"data"`
	Logging *vlogger.LogConfig `yaml:"logging" json:"logging"`
}

type Server struct {
	Http *vbasedata.Http `yaml:"http" json:"http"`
}

type Data struct {
	Redis *vbasedata.RedisConfig `yaml:"redis" json:"redis"`
	DB    *vbasedata.GormConfig  `yaml:"db" json:"db"`
	Kafka *vbasedata.KafkaConfig `yaml:"kafka" json:"kafka"`
	Pond  *vbasedata.Pond        `yaml:"pond" json:"pond"`
}
