package service

import (
	"net/http"
	"time"
	"vmapp/app0/internal/dto"
	"vmapp/pkg/vhttp"

	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kataras/iris/v12"
)

type SysService struct {
	c  *conf.AppConf
	uc *usecase.UserUseCase
	bc *conf.BootComponent

	log *log.Helper
}

func NewSysService(c *conf.AppConf, bc *conf.BootComponent, uc *usecase.UserUseCase) *SysService {
	return &SysService{
		c:   c,
		uc:  uc,
		bc:  bc,
		log: bc.Logger,
	}
}

func (t *SysService) Login(ctx *gin.Context) {
	var req dto.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := t.uc.LoginUser(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, res)
}

func (t *SysService) LogOut(ctx *gin.Context) {
	token := ctx.Request.Header.Get("VToken")

	err := t.uc.LogOutUser(ctx, token)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}

	vhttp.SuccessHandle(ctx, nil)
}

func (t *SysService) ReferToken(ctx *gin.Context) {
	token := ctx.Request.Header.Get("VToken")
	if token == "" {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, "token不存在"))
		return
	}

	res, err := t.uc.ReferTokenUser(ctx, token)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, res)
}

func (t *SysService) Register(ctx *gin.Context) {
	var req dto.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	err := t.uc.Register(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *SysService) Forget(ctx *gin.Context) {
	var req dto.ForgetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	err := t.uc.Forget(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *SysService) UpdateUserPass(ctx *gin.Context) {
	var req dto.UpdateUserPassReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	err := t.uc.UpdateUserPass(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *SysService) UpdateUserProfile(ctx *gin.Context) {
	var req dto.UpdateUserProfileReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	err := t.uc.UpdateUserProfile(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}

func (t *SysService) Captcha(ctx *gin.Context) {
	id, imgbase64, _, err := t.bc.Captcha.GetCaptCha(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("验证码生成失败: %v", err)
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, "验证码生成失败", vhttp.WithReason(err)))
		return
	}
	vhttp.SuccessHandle(ctx, iris.Map{
		"id":        id,
		"imgbase64": imgbase64,
	})

}

func (t *SysService) Restart(ctx *gin.Context) {
	vhttp.SuccessHandle(ctx, nil)
	time.Sleep(1 * time.Second)
	go func() {
		t.bc.RestartChan <- 1
	}()
}
