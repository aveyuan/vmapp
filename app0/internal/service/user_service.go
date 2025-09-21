package service

import (
	"net/http"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/app0/internal/middleware"
	"vmapp/pkg/vhttp"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	bcf *conf.Data
	uc  *usecase.UserUseCase
	bc  *conf.BootComponent
}

func NewUserService(bcf *conf.Data, bc *conf.BootComponent, uc *usecase.UserUseCase) *UserService {
	return &UserService{
		bcf: bcf,
		uc:  uc,
		bc:  bc,
	}
}

func (t *UserService) DeleteUser(ctx *gin.Context) {
	var req dto.DeleteUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.uc.DelectUser(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *UserService) UpdateUser(ctx *gin.Context) {
	var req dto.UpdateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.uc.UpdateUser(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *UserService) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.uc.CreateUser(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *UserService) ListUser(ctx *gin.Context) {
	var req dto.ListUserReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	count, all, err := t.uc.ListUser(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, all, vhttp.WithPage(vhttp.NewPaginator(int(count), req.Page, req.Limit)))
}

func (t *UserService) GetUser(ctx *gin.Context) {
	var req dto.GetUserReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	one, err := t.uc.GetUser(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, one)
}

func (t *UserService) RepassUser(ctx *gin.Context) {
	var req dto.RepassUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.uc.RepassUser(ctx, &req); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *UserService) ResetToken(ctx *gin.Context) {

	if err := t.uc.ReSetToekn(ctx); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *UserService) GetUserInfo(ctx *gin.Context) {

	userp, err := middleware.GetCtxUser(ctx)
	if err != nil {

		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusInternalServerError, "获取用户信息失败", vhttp.WithReason(err)))
		return
	}

	user, err := t.uc.GetUser(ctx, &dto.GetUserReq{UID: userp.Uid})
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, user)
}
