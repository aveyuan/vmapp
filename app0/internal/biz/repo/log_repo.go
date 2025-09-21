package repo

import (
	"context"
	"vmapp/app0/internal/dto"
	"vmapp/app0/internal/models"
)

type EventLogType int8
type EventLogModel int8

const (
	EventLogTypeAdd    EventLogType = 1
	EventLogTypeEdit   EventLogType = 2
	EventLogTypeDelete EventLogType = 3
	EventLogTypeQuery  EventLogType = 4
	EventLogTypeOp     EventLogType = 5
)

var EventLogTypeMap = map[EventLogType]string{
	EventLogTypeAdd:    "新增",
	EventLogTypeEdit:   "编辑",
	EventLogTypeDelete: "删除",
	EventLogTypeQuery:  "查询",
	EventLogTypeOp:     "操作",
}

const (
	EventLogAdminSystem EventLogModel = 1
)

var EventLogModelMap = map[EventLogModel]string{
	EventLogAdminSystem: "管理端/系统管理",
}

type EventLog struct {
	Type         EventLogType  `json:"type" form:"type"`                   //类型
	Model        EventLogModel `json:"model" form:"model"`                 //模块
	ObjId        interface{}   `json:"obj_id" form:"obj_id"`               //对象ID
	ObjName      string        `json:"obj_name" form:"obj_name"`           //对象名称
	Op           string        `json:"op" form:"op"`                       //操作行为
	ChangeName   string        `json:"change_name" form:"change_name"`     //变更名称
	ChangeBefore string        `json:"change_before" form:"change_before"` //变更前
	ChangeAfer   string        `json:"change_afer" form:"change_afer"`     //变更后
}

type EventLogRepo interface {
	CreateEventLog(ctx context.Context, one *models.EventLog) error
	ListEventLog(ctx context.Context, req *dto.ListEventLogReq) (count int64, all []*models.EventLog, err error)
	ListSendLog(ctx context.Context, req *dto.ListSendLogReq) (count int64, all []*models.Send, err error)
}
