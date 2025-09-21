package models

import "time"

// EventLog 事件日志表
type EventLog struct {
	Id          int32     `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:业务ID" json:"id" form:"id"`
	Type         int8      `gorm:"column:type;type:tinyint(4);comment:操作类型" json:"type" form:"type"`
	Model        int8      `gorm:"column:model;type:tinyint(4);comment:模块" json:"model" form:"model"`
	Uid          string    `gorm:"column:uid;type:varchar(20);comment:用户ID" json:"uid" form:"uid"`
	Uaccount     string    `gorm:"column:uaccount;type:varchar(50);comment:用户账户" json:"uaccount" form:"uaccount"`
	Uname        string    `gorm:"column:uname;type:varchar(50);comment:用户名称" json:"uname" form:"uname"`
	ObjId        string    `gorm:"column:obj_id;type:varchar(40);comment:操作对象ID" json:"obj_id" form:"obj_id"`
	ObjName      string    `gorm:"column:obj_name;type:varchar(100);comment:操作对象名称" json:"obj_name" form:"obj_name"`
	Op           string    `gorm:"column:op;type:varchar(255);comment:操作行为" json:"op" form:"op"`
	ChangeName   string    `gorm:"column:change_name;type:varchar(20);comment:变更名称" json:"change_name" form:"change_name"`
	ChangeBefore string    `gorm:"column:change_before;type:varchar(200);comment:变更前内容" json:"change_before" form:"change_before"`
	ChangeAfer   string    `gorm:"column:change_afer;type:varchar(200);comment:变更后内容" json:"change_afer" form:"change_afer"`
	Ip           string    `gorm:"column:ip;type:varchar(20);comment:IP地址" json:"ip" form:"ip"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"created_at" form:"created_at"`
}
