package send_repo

import (
	"context"
	"errors"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/vconst"

	"github.com/go-kratos/kratos/v2/log"
)

type sendRepo struct {
	data *base.Data
	log  *log.Helper
	bc   *conf.BootComponent
}

func NewSendRepo(data *base.Data, component *conf.BootComponent) repo.SendRepo {
	return &sendRepo{
		data: data,
		log:  component.Logger,
		bc:   component,
	}
}

// if c.Data.Email != nil {
// 	data.Email = basedata.NewEmail(c.Data.Email)
// }

func (t *sendRepo) SendMsg(ctx context.Context, media vconst.SendMedia, sendType vconst.SendType, Template string, to string, title string, msg string) error {
	return errors.New("not impl")
}
