package conf

import (
	"time"

	"github.com/aveyuan/vbasedata"
	"github.com/aveyuan/vjwt"
	"github.com/go-kratos/kratos/v2/log"
)

type VUser struct {
	Uid     int64    `json:"uid" form:"uid"`
	Uname   string   `json:"uname" form:"uname"`
	Role    []string `json:"role" form:"role"`
	IsAdmin bool     `json:"is_admin" form:"is_admin"`
}

// BootComponent 用于全局配置
type BootComponent struct {
	Logger      *log.Helper
	Pond        *vbasedata.Pond
	Captcha     *vbasedata.Captcha
	Idgenerator *vbasedata.Idgenerator
	Jwt         *vjwt.Vjwt[VUser]
	LruCode     *vbasedata.LruCache //验证码缓存
	LruLock     *vbasedata.LruCache //验证码获取锁定缓存
	LoginCount  *vbasedata.LruCache //登录计数
	LoginLock   *vbasedata.LruCache //登录锁定
	RestartChan chan int
}

// NewBootComponent 初始化
func NewBootComponent(dc *AppConf, logger *log.Helper) *BootComponent {
	bc := &BootComponent{
		Logger:      logger,
		Captcha:     vbasedata.NewCaptcha(&vbasedata.CaptchaConfig{}, vbasedata.NewLruCache(500, 300*time.Second)),
		RestartChan: make(chan int),
		Idgenerator: vbasedata.NewIdgenerator(0),
		LruCode:     vbasedata.NewLruCache(200, 5*time.Minute),
		LoginCount:  vbasedata.NewLruCache(200, 3*time.Minute),
		LoginLock:   vbasedata.NewLruCache(200, 5*time.Minute),
		LruLock:     vbasedata.NewLruCache(200, 1*time.Minute),
	}

	if dc.Data.Pond != nil {
		bc.Pond = vbasedata.NewPond(dc.Data.Pond, logger)
	} else {
		bc.Pond = vbasedata.NewPond(&vbasedata.PondConfig{MaxWorkers: 2, MinWorkers: 2, MaxCapacity: 0, StopAndWait: 5}, logger)
	}

	if dc.Data.Jwt != nil {
		bc.Jwt = vjwt.NewJwt[VUser](dc.Data.Jwt)
	}

	return bc
}
