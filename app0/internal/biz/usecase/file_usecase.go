package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"
	"vmapp/pkg/vhttp"

	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/app0/internal/middleware"
	"vmapp/app0/internal/models"
	"vmapp/pkg/encrypt/sha"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var ImgList []string = []string{".jpg", ".png", ".jpeg", ".gif", ".icon", ".webp", ".psd"}
var DocList []string = []string{".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".md"}

type FileUseCase struct {
	fp  repo.FileRepo
	log *log.Helper
	bc  *conf.BootComponent
}

type FileRes struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func NewFileUseCase(fp repo.FileRepo, tx repo.Transaction, bc *conf.BootComponent) *FileUseCase {
	return &FileUseCase{
		fp:  fp,
		log: bc.Logger,
		bc:  bc,
	}
}

func (t *FileUseCase) UploadFile(ctx context.Context, Exts []string) (*FileRes, error) {
	fileRes := &FileRes{}
	user, err := middleware.GetCtxUser(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取用户信息失败,%v", err)
		return fileRes, vhttp.NewError(http.StatusInternalServerError, "获取用户信息失败", vhttp.WithReason(err))
	}

	f, err := ctx.(*gin.Context).FormFile("file")
	if err != nil {
		return fileRes, vhttp.NewError(http.StatusBadRequest, "文件解析失败", vhttp.WithReason(err))
	}
	file, err := f.Open()
	if err != nil {
		t.log.WithContext(ctx).Errorf("文件打开失败,%v", err)
		return fileRes, vhttp.NewError(http.StatusBadRequest, "文件打开失败", vhttp.WithReason(err))
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(f.Filename))

	if len(Exts) > 0 {
		var match bool = false

		for _, v := range Exts {
			if v == strings.ToLower(ext) {
				match = true
				break
			}
		}

		if !match {
			return fileRes, vhttp.NewError(http.StatusBadRequest, "当前文件名上传不被允许")
		}

	}

	// if file.Size > size {
	// 	return fileRes, vhttp.NewError(http.StatusBadRequest, "当前文件大小超过限制")
	// }

	// 查看hash
	// shaOpen, err := file.Open()
	// if err != nil {
	// 	t.log.WithContext(ctx).Errorf("文件打开失败,%v", err)
	// 	return fileRes, vhttp.NewError(http.StatusBadRequest, "文件打开失败", vhttp.WithReason(err))
	// }

	// defer shaOpen.Close()
	shaByte, err := io.ReadAll(file)
	if err != nil {
		t.log.WithContext(ctx).Errorf("文件读取失败,%v", err)
		return fileRes, vhttp.NewError(http.StatusBadRequest, "文件读取失败", vhttp.WithReason(err))
	}

	sha256 := sha.Sha256Byte(shaByte)
	res, err := t.fp.GetFileByHash(ctx, sha256)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.WithContext(ctx).Errorf("文件数据查询失败,%v", err)
		return fileRes, vhttp.NewError(http.StatusBadRequest, "文件数据查询失败", vhttp.WithReason(err))
	}

	if res != nil && res.Id != 0 {
		fileRes.Name = res.Name
		fileRes.Path = res.Path
		return fileRes, nil
	}

	dir := fmt.Sprintf("/%v", time.Now().Format("2006/01/02"))
	id := t.bc.Idgenerator.NextId()
	filename := fmt.Sprintf("%v%v", id, ext)
	localDir := "/uploads"

	fileModel := &models.File{
		Id:        id,
		Size:      f.Size,
		Name:      f.Filename,
		Sha256:    sha256,
		Suffix:    ext,
		Mime:      f.Header.Get("Content-Type"),
		Uid:       user.Uid,
		CreatedAt: time.Now(),
		Private:   0,
	}

	fileModel.Type = func() int8 {
		for _, v := range ImgList {
			if ext == v {
				return 2
			}
		}
		for _, v := range DocList {
			if ext == v {
				return 3
			}
		}
		return 1
	}()



	vkey, err := t.fp.UploadFile(ctx, repo.Vkey{
		Dir:      dir,
		FileName: filename,
		LocalDir: localDir,
	}, file)
	if err != nil {
		t.log.WithContext(ctx).Errorf("文件存储失败,%v", err)
		return fileRes, vhttp.NewError(http.StatusInternalServerError, "文件存储失败", vhttp.WithReason(err))
	}

	fileModel.Path = vkey
	fileModel.Media = 0

	if err := t.fp.CreateFile(ctx, fileModel); err != nil {
		// 上传成功，记录失败，不阻塞
		t.log.WithContext(ctx).Errorf("文件表写入失败,%v", err)
	}
	fileRes.Name = fileModel.Name
	fileRes.Path = fileModel.Path
	return fileRes, nil
}

func (t *FileUseCase) DeleteFile(ctx context.Context, FileID int64) error {
	one, err := t.fp.GetFile(ctx, FileID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.WithContext(ctx).Errorf("获取文件失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "获取文件失败", vhttp.WithReason(err))
	}
	if one == nil {
		return vhttp.NewError(http.StatusBadRequest, "文件不存在", vhttp.WithReason(err))
	}

	// 删除
	if err := t.fp.DeleteUploadFile(ctx, int(one.Media), one.Path); err != nil {
		t.log.WithContext(ctx).Errorf("文件删除失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "文件删除失败", vhttp.WithReason(err))
	}

	if err := t.fp.DeleteFile(ctx, FileID); err != nil {
		t.log.WithContext(ctx).Errorf("文件删除失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "文件删除失败", vhttp.WithReason(err))
	}
	return nil
}

func (t *FileUseCase) ListFile(ctx context.Context, req *dto.ListFileReq) (count int64, all []*models.File, err error) {
	count, all, err = t.fp.ListFile(ctx, req)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取文件列表失败,%v", err)
		return 0, nil, vhttp.NewError(http.StatusInternalServerError, "获取文件列表失败", vhttp.WithReason(err))
	}
	return
}
