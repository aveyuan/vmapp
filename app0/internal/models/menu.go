package models

import "time"

// Menu 菜单表
type Menu struct {
	Id          int32     `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;comment:菜单ID" json:"id" form:"id"`
	Pid         int32     `gorm:"column:pid;type:int;comment:上级ID" json:"pid" form:"pid"`
	Title       string    `gorm:"column:title;type:varchar(30);comment:名称" json:"title" form:"title"`
	Type        int8      `gorm:"column:type;type:tinyint;not null;comment:类型  1 菜单  2按钮" json:"type" form:"type"`
	Icon        string    `gorm:"column:icon;type:varchar(30);comment:icon标识" json:"icon" form:"icon"`
	Href        string    `gorm:"column:href;type:varchar(255);comment:打开方法" json:"href" form:"href"`
	Target      string    `gorm:"column:target;type:varchar(10);comment:资源地址" json:"target" form:"target"`
	ApiMethod   string    `gorm:"column:api_method;type:varchar(10);comment:接口方法" json:"api_method" form:"api_method"`
	ApiUrl      string    `gorm:"column:api_url;type:varchar(100);comment:接口地址" json:"api_url" form:"api_url"`
	ApiAuthCode string    `gorm:"column:api_auth_code;type:varchar(30);comment:权限标识" json:"api_auth_code" form:"api_auth_code"`
	Sort        int8      `gorm:"column:sort;type:tinyint;index;not null;comment:排序" json:"sort" form:"sort"`
	Remark      string    `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark" form:"remark"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:更新时间" json:"updated_at" form:"updated_at"`
}
