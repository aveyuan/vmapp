package models

import (
	"time"
)

// Send 消息发送表
type Send struct {
	Id        int32     `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:业务ID" json:"id" form:"id"`
	Media     int8      `gorm:"column:meda;type:int; not null;comment:媒体介质 1邮箱 2手机" json:"type" form:"type"`
	SendType  int8      `gorm:"column:sned_type;type:int; not null;comment:发送类型 1登录 2注册 3找回" json:"send_type" form:"send_type"`
	To        string    `gorm:"column:to;type:varchar(100);comment:发送到" json:"to" form:"to"`
	Title     string    `gorm:"column:title;type:varchar(50);comment:消息标题" json:"title" form:"title"`
	Msg       string    `gorm:"column:to;type:text;comment:消息内容" json:"msg" form:"msg"`
	Template  string    `gorm:"column:template;type:varchar(20);comment:消息模板" json:"template" form:"template"`
	Meta      string    `gorm:"column:meta;type:varchar(255);comment:附加信息" json:"meta" form:"meta"`
	Ip        string    `gorm:"column:ip;type:varchar(20);comment:IP地址" json:"ip" form:"ip"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"created_at" form:"created_at"`
}
