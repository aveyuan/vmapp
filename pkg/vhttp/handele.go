package vhttp

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
)

const ErrCtx = "errCtx"

func SuccessHandle(ctx context.Context, data interface{}) {

	// 如果是gin，则使用gin渲染
	if c, ok := ctx.(*gin.Context); ok {
		c.JSON(http.StatusOK, data)
	}

	// 如果是Iris，则使用Iris渲染
	if c, ok := ctx.(iris.Context); ok {
		c.StatusCode(http.StatusOK)
		c.JSON(data)
	}

}

func ErrorHandle(ctx context.Context, err error) {
	var errData ErrData
	if !errors.As(err, &errData) {
		errData.Code = http.StatusInternalServerError
		errData.Msg = "系统错误"
		errData.Reason = err
	}
	// 如果是gin，则使用gin渲染
	if c, ok := ctx.(*gin.Context); ok {
		c.Set(ErrCtx, &errData)
		if errData.StatusCode != nil {
			c.JSON(*errData.StatusCode, errData)
			return
		}
		c.JSON(200, errData)
		return

	}

	// 如果是iris，则使用iris渲染
	if c, ok := ctx.(iris.Context); ok {
		c.Values().Set(ErrCtx, errData)
		if errData.StatusCode != nil {
			c.StatusCode(*errData.StatusCode)
		}
		c.JSON(errData)
	}

}
