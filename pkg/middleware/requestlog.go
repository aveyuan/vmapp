package middleware

import (
	"time"

	"github.com/aveyuan/vlogger"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Ext struct {
	UID           string  `json:"uid,omitempty"`
	Uname         string  `json:"uname,omitempty"`
	Referer       string  `json:"referer,omitempty"`
	Path          string  `json:"path,omitempty"`
	Query         string  `json:"query,omitempty"`
	ClientIP      string  `json:"client_ip,omitempty"`
	Method        string  `json:"method,omitempty"`
	ReqTime       string  `json:"req_time,omitempty"`
	RespTime      string  `json:"resp_time,omitempty"`
	TotalTimeUsed float64 `json:"total_time_used,omitempty"`
	StatusCode    int     `json:"status_code,omitempty"`
	UserAgent     string  `json:"user_agent,omitempty"`
}

// NewGinLog Gin访问日志
func NewRequestLog(logger *log.Helper) ghttp.HandlerFunc {
	return func (ctx *ghttp.Request) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.Query().Encode()
		ctx.Middleware.Next()

		end := time.Now()
		latency := end.Sub(start).Seconds()

		// 判断是否有错误返回
		ext := &Ext{
			StatusCode:    ctx.Response.Status,
			ReqTime:       start.Format(time.DateTime),
			RespTime:      end.Format(time.DateTime),
			Path:          path,
			Method:        ctx.Request.Method,
			Query:         query,
			ClientIP:      ctx.GetClientIp(),
			UserAgent:     ctx.Request.UserAgent(),
			TotalTimeUsed: latency,
		}
		// 补充错误信息
		
		if ctx.Response.Status >= 500 {
			logger.WithContext(vlogger.WithExt(ctx.Context(), &vlogger.ExtLogValue{Ext: ext})).Error()
		} else {
			logger.WithContext(vlogger.WithExt(ctx.Context(), &vlogger.ExtLogValue{Ext: ext})).Info()
		}

	}
}
