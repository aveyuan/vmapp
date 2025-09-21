package repo

import (
	"context"
	"io"
	"vmapp/app0/internal/dto"
	"vmapp/app0/internal/models"
)

type Vkey struct {
	Dir string
	FileName string
	LocalDir string
}

type FileRepo interface {
	CreateFile(ctx context.Context, one *models.File) error
	GetFileByHash(ctx context.Context, hash string) (one *models.File, err error)
	GetFile(ctx context.Context, id int64) (one *models.File, err error)
	ListFile(ctx context.Context, req *dto.ListFileReq) (count int64, all []*models.File, err error)
	DeleteFile(ctx context.Context, id int64) (err error)
	DeleteUploadFile(ctx context.Context,media int, vkey string) (error)
	UploadFile(ctx context.Context, vkey Vkey, r io.Reader) (string, error)
}
