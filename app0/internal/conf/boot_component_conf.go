package conf

import (


	"github.com/alitto/pond"
	"github.com/go-kratos/kratos/v2/log"
)

// BootComponent 用于全局配置
type BootComponent struct {
	Logger *log.Helper
	Pond   *pond.WorkerPool
}

// NewBootComponent 初始化
func NewBootComponent(logger *log.Helper) *BootComponent {
	return &BootComponent{
		Logger: logger,
	}
}
