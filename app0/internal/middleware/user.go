package middleware

import (
	"context"
	"errors"
	"vmapp/app0/internal/conf"

	"github.com/gin-gonic/gin"
)

// GetCtxUser 获取用户信息
func GetCtxUser(ctx context.Context) (*conf.VUser, error) {
	
	if ctx, ok := ctx.(*gin.Context); ok {
		vk, ok := ctx.Get(UserContextKey)
		if !ok {
			return nil, errors.New("获取用户信息失败")
		}
		if u, ok := vk.(conf.VUser); ok {
			return &u, nil
		}
	}
	return nil, errors.New("获取用户信息失败")
}
