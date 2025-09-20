package vhttp

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ErrCtx = "errCtx"

func SuccessHandle(ctx *gin.Context, data interface{}, opts ...DataOptions) {
	d := &Data{Code: 200, Msg: "ok", Data: data}
	for _, v := range opts {
		v(d)
	}
	ctx.JSON(d.Code, d)
}

func ErrorHandle(ctx *gin.Context, err error) {
	var errData ErrData
	if !errors.As(err, &errData) {
		errData.Code = http.StatusInternalServerError
		errData.Msg = "系统错误"
		errData.Reason = err
	}
	ctx.Set(ErrCtx, &errData)
	if errData.StatusCode != nil {
		ctx.JSON(*errData.StatusCode, errData)
		return
	}
	ctx.JSON(200, errData)
}
