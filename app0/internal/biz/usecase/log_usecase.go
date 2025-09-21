package usecase

import (
	"context"
	"net/http"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/app0/internal/models"
	"vmapp/app0/internal/vconst"
	"vmapp/pkg/vhttp"

	"github.com/go-kratos/kratos/v2/log"
)

type LogUseCase struct {
	c   *conf.Data
	bc  *conf.BootComponent
	log *log.Helper

	ep repo.EventLogRepo
}

func NewLogUseCase(c *conf.Data, bc *conf.BootComponent, ep repo.EventLogRepo) *LogUseCase {
	return &LogUseCase{
		c:   c,
		log: bc.Logger,
		bc:  bc,
		ep:  ep,
	}
}

type ListEvent struct {
	*models.EventLog
	TypeStr  string `json:"type_str"`
	ModelStr string `json:"model_str"`
}

type ListSend struct {
	*models.Send
	MediaStr    string `json:"media_str"`
	SendTypeStr string `json:"send_type_str"`
}

func (t *LogUseCase) ListEventLog(ctx context.Context, req *dto.ListEventLogReq) (count int64, all []*ListEvent, err error) {
	count, res, err := t.ep.ListEventLog(ctx, req)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取日志列表失败,%v", err)
		return 0, nil, vhttp.NewError(http.StatusInternalServerError, "获取日志列表失败", vhttp.WithReason(err))
	}
	for _, v := range res {
		all = append(all, &ListEvent{EventLog: v, TypeStr: repo.EventLogTypeMap[repo.EventLogType(v.Type)], ModelStr: repo.EventLogModelMap[repo.EventLogModel(v.Model)]})
	}
	return count, all, nil
}

func (t *LogUseCase) ListSendLog(ctx context.Context, req *dto.ListSendLogReq) (count int64, all []*ListSend, err error) {
	count, res, err := t.ep.ListSendLog(ctx, req)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取日志列表失败,%v", err)
		return 0, nil, vhttp.NewError(http.StatusInternalServerError, "获取日志列表失败", vhttp.WithReason(err))
	}
	for _, v := range res {
		all = append(all, &ListSend{Send: v, MediaStr: vconst.SendMediaMap[vconst.SendMedia(v.Media)], SendTypeStr: vconst.SendTypeMap[vconst.SendType(v.SendType)]})
	}
	return count, all, nil
}
