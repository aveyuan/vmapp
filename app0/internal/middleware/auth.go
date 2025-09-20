package middleware

import (
	"net/http"
	"vmapp/app0/internal/conf"

	"github.com/gin-gonic/gin"
)

const TokenKey = "VToken"
const UserContextKey = "VUserContext"

// Auth 认证
func Auth(bc *conf.BootComponent) gin.HandlerFunc {
	return func(ctx *gin.Context) {

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
		ctx.Next()
	}
}
