package send_repo

import (
	"context"
	"errors"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"
	"vmapp/app0/internal/vconst"

	"github.com/aveyuan/vbasedata"
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

func (t *sendRepo) SendMsg(ctx context.Context, media vconst.SendMedia, sendType vconst.SendType, Template string, to string, title string, msg string) error {

	if t.bc.Email == nil {
		return errors.New("当前邮箱未配置")
	}

	if err := t.bc.Email.SendMsg(&vbasedata.Msg{
		Title:    title,
		Body:     msg,
		To:       to,
		BodyType: vbasedata.TextBodyType,
	}); err != nil {
		return err
	}

	return nil
}
