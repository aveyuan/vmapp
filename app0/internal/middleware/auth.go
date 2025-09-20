package middleware

import (
	"net/http"
	"strings"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"

	"github.com/gin-gonic/gin"
)

const TokenKey = "VToken"
const UserContextKey = "VUserContext"

// Auth 认证
func Auth(bc *conf.BootComponent, data *base.Data) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 不要登录的
		var expBaseList = []string{"/login", "/register", "/forget", "/logout", "/static", "/captcha", "/sys", "/getcode"}
		// 登录后不做鉴权的
		var expLoginList = []string{"/api/menu/system_menu", "/api/sys/"}

		var notValid bool
		path := ctx.Request.URL.Path
		for _, v := range expBaseList {
			if strings.HasPrefix(path, v) {
				notValid = true
				break
			}
		}
		if notValid {
			ctx.Next()
			return
		}

		token := ctx.Request.Header.Get(TokenKey)
		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"msg": "未登录"})
			return
		}
		jc, err := bc.Jwt.Verify(token)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"msg": "请重新登录"})
			return
		}

		var user conf.VUser
		if err := jc.Claims(&user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "用户解析失败"})
			return
		}

		ctx.Set(UserContextKey, user)

		for _, v := range expLoginList {
			if strings.HasPrefix(path, v) {
				notValid = true
				break
			}
		}
		if notValid {
			ctx.Next()
			return
		}

		for _, v := range user.Role {
			// 得到权限,进行检查
			if data.Rbac.Enforce(v, path, ctx.Request.Method) {
				ctx.Next()
				return
			}
		}

		// 没有通过权限检查，返回
		ctx.AbortWithStatusJSON(401, gin.H{"msg": "权限不足"})
		return

	}
}
