package repo

import (
	"context"
	"vmapp/app0/internal/vconst"
)



type SendRepo interface {
	SendMsg(ctx context.Context, media vconst.SendMedia, sendType vconst.SendType, Template string, to string, title string, msg string) error
}
