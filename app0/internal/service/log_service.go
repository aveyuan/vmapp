package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/pkg/vhttp"
)

type LogService struct {
	c  *conf.Data
	ec *usecase.LogUseCase
	bc *conf.BootComponent
}

func NewLogService(c *conf.Data, bc *conf.BootComponent, ec *usecase.LogUseCase) *LogService {
	return &LogService{
		c:  c,
		ec: ec,
		bc: bc,
	}
}

func (t *LogService) ListEventLog(ctx *gin.Context) {
	var req dto.ListEventLogReq
	if err := ctx.ShouldBind(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	count, all, err := t.ec.ListEventLog(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, all, vhttp.WithPage(vhttp.NewPaginator(int(count), req.Page, req.Limit)))
}

func (t *LogService) ListSendLog(ctx *gin.Context) {
	var req dto.ListSendLogReq
	if err := ctx.ShouldBind(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest,err.Error()))
		return
	}
	count, all, err := t.ec.ListSendLog(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, all, vhttp.WithPage(vhttp.NewPaginator(int(count), req.Page, req.Limit)))
}

