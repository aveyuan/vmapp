package service

import (
	"net/http"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/pkg/vhttp"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type RoleService struct {
	c  *conf.Data
	mc *usecase.RoleUseCase
	bc *conf.BootComponent
}

func NewRoleService(c *conf.Data, bc *conf.BootComponent, mc *usecase.RoleUseCase) *RoleService {
	return &RoleService{
		c:  c,
		mc: mc,
		bc: bc,
	}
}

func (t *RoleService) DeleteRole(ctx *gin.Context) {
	var req dto.DeleteRoleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.mc.DelectRole(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *RoleService) UpdateRole(ctx *gin.Context) {
	var req dto.UpdateRoleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest,err.Error()))
		return
	}
	
	if err := t.mc.UpdateRole(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *RoleService) CreateRole(ctx *gin.Context) {
	var req dto.CreateRoleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest,err.Error()))
		return
	}
	
	if err := t.mc.CreateRole(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *RoleService) GetUserRole(ctx *gin.Context) {
	uidstr := ctx.Query("uid")
	uid, err := cast.ToInt64E(uidstr)
	if err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))	
		return
	}

	one, err := t.mc.GetUserRole(ctx, uid)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, one)
}

func (t *RoleService) ListRole(ctx *gin.Context) {
	all, err := t.mc.ListRole(ctx)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, all)
}

func (t *RoleService) GetRole(ctx *gin.Context) {
	var req dto.GetRoleReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	
	one, err := t.mc.GetRole(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, one)
}

