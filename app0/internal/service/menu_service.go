package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/pkg/vhttp"
)

type MenuService struct {
	c  *conf.Data
	mc *usecase.MenuUseCase
	bc *conf.BootComponent
}

func NewMenuService(c *conf.Data, bc *conf.BootComponent, mc *usecase.MenuUseCase) *MenuService {
	return &MenuService{
		c:  c,
		mc: mc,
		bc: bc,
	}
}

func (t *MenuService) DeleteMenu(ctx *gin.Context) {
	var req dto.DeleteMenuReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.mc.DelectMenu(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *MenuService) UpdateMenu(ctx *gin.Context) {
	var req dto.UpdateMenuReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.mc.UpdateMenu(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *MenuService) CreateMenu(ctx *gin.Context) {
	var req dto.CreateMenuReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.mc.CreateMenu(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *MenuService) ListMenu(ctx *gin.Context) {
	all, err := t.mc.ListMenu(ctx)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, all)
}
func (t *MenuService) ListMenuTree(ctx *gin.Context) {
	all, err := t.mc.ListMenuTree(ctx)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, all)
}

func (t *MenuService) GetMenu(ctx *gin.Context) {
	var req dto.GetMenuReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	one, err := t.mc.GetMenu(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, one)
}
