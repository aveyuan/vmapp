package models

import "time"

// Role 角色表
type Role struct {
	Id        int32     `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:业务ID" json:"id" form:"id"`
	RoleName  string    `gorm:"column:role_name;type:varchar(20);comment:权限名称" json:"role_name" form:"role_name"`
	RoleCode  string    `gorm:"column:role_code;type:varchar(20);uniqueIndex;comment:权限标识" json:"role_code" form:"role_code"`
	MenuIds   string    `gorm:"column:menu_ids;type:text;comment:关联的菜单ID" json:"menu_ids" form:"menu_ids"`
	IndexMenu string    `gorm:"column:index_menu;type:varchar(20);comment:首页菜单" json:"index_menu" form:"index_menu"`
	Remark    string    `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark" form:"remark"`
	UpdateUid int64     `gorm:"column:update_uid;type:bigint;comment:更新用户" json:"update_uid" form:"update_uid"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:更新时间" json:"updated_at" form:"updated_at"`
}
