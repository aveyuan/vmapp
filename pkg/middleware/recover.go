package middleware

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NewRecoverLog(logger *log.Helper) ghttp.HandlerFunc {
	return func (r *ghttp.Request) {
		r.Middleware.Next()
		if err := r.GetError(); err != nil {
			logger.WithContext(r.GetCtx()).Errorf("%+v", err)
			r.Response.Status = 500
			r.Response.Write("Internal Server Error")
		}
	}
}