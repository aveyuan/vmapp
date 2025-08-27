package conf

import (
	"github.com/aveyuan/vbasedata"

	"github.com/aveyuan/vlogger"

)

type AppConf struct {
	App     *vbasedata.App      `yaml:"app" json:"app"`
	Server  *Server            `yaml:"server" json:"server"`
	Data    *Data              `yaml:"data" json:"data"`
	Logging *vlogger.LogConfig `json:"logging"`
}

type Server struct {
	Http *vbasedata.Http `json:"http"`
}

type Data struct {
	Redis *vbasedata.RedisConfig `json:"redis"`
	DB    *vbasedata.GormConfig  `json:"db"`
	Kafka *vbasedata.KafkaConfig `json:"kafka"`
	Pond  *vbasedata.Pond        `json:"pond"`
}
