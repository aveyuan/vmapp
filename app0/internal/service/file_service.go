package service

import (
	"net/http"
	"vmapp/app0/internal/biz/usecase"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/pkg/vhttp"

	"github.com/gin-gonic/gin"
)

type FileService struct {
	c  *conf.Data
	fc *usecase.FileUseCase
	bc *conf.BootComponent
}

func NewFileService(c *conf.Data, bc *conf.BootComponent, fc *usecase.FileUseCase) *FileService {
	return &FileService{
		c:  c,
		fc: fc,
		bc: bc,
	}
}

func (t *FileService) UploadFile(ctx *gin.Context) {
	res, err := t.fc.UploadFile(ctx, nil)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, res)
}

func (t *FileService) ListFile(ctx *gin.Context) {
	var req dto.ListFileReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	count, all, err := t.fc.ListFile(ctx, &req)
	if err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, all, vhttp.WithPage(vhttp.NewPaginator(int(count), req.Page, req.Limit)))
}

func (t *FileService) DeleteFile(ctx *gin.Context) {
	var req dto.DeleteFileReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		vhttp.ErrorHandle(ctx, vhttp.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := t.fc.DeleteFile(ctx, req.ID); err != nil {
		vhttp.ErrorHandle(ctx, err)
		return
	}
	vhttp.SuccessHandle(ctx, nil)
}
