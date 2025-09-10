package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	Uid          int64          `gorm:"column:uid;type:bigint;not null;primaryKey;autoIncrement;comment:用户ID" json:"uid" form:"uid"`
	Username     string         `gorm:"column:username;type:varchar(32);uniqueIndex;comment:用户名" json:"username" form:"username"`
	Nickname     string         `gorm:"column:nickname;type:varchar(32);comment:名称" json:"nickname" form:"nickname"`
	Avatar       string         `gorm:"column:avatar;type:varchar(50);comment:头像" json:"avatar" form:"avatar"`
	Email        string         `gorm:"column:email;type:varchar(150);index;comment:邮箱" json:"email" form:"email"`
	Url          string         `gorm:"column:url;type:varchar(100);comment:个人地址" json:"url" form:"url"`
	Description  string         `gorm:"column:description;type:VARCHAR(255);comment:个人简介" json:"description" form:"description"`
	RoleCodes    string         `gorm:"column:role_codes;type:varchar(64);comment:权限标识" json:"role_codes" form:"role_codes"`
	Password     string         `gorm:"column:password;type:varchar(64);comment:密码" json:"password" form:"password"`
	Salt         string         `gorm:"column:salt;type:varchar(10);comment:加盐" json:"salt" form:"salt"`
	Template     string         `gorm:"column:template; type:VARCHAR(50);default:;comment:模板路径" json:"template" form:"template"`
	LastLogin    time.Time      `gorm:"column:last_login;type:datetime;comment:最后登录时间" json:"last_login" form:"last_login"`
	Token        string         `gorm:"column:token;type:VARCHAR(255);comment:token" json:"token" form:"token"`
	CheckStatus  int8           `gorm:"column:check_status;type:tinyint;default:-1;not null;comment:审核状态 -1 待审核 1 审核通过 2 审核拒绝" json:"check_status" form:"check_status"`
	Status       int8           `gorm:"column:status;type:tinyint;default:-1;not null;comment:状态 1 启用 2 禁用" json:"status" form:"status"`
	ArticleCount int32          `gorm:"column:article_count;type:int;default:0;not null;comment:文章数量" json:"article_count" form:"article_count"`
	IsAdmin      int8           `gorm:"column:is_admin;type:tinyint;default:2;not null;comment:是否是管理员 1是 2 不是" json:"is_admin" form:"is_admin"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"created_at" form:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:更新时间" json:"updated_at" form:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index;comment:删除标识" json:"deleted_at" form:"deleted_at"`
}
