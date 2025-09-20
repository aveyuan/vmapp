package models

import "time"

// File 文件表
type File struct {
	Id        int64     `gorm:"column:id;type:bigint;not null;primaryKey;autoIncrement;comment:业务ID" json:"id" form:"id"`
	Type      int8      `gorm:"column:type;type:tinyint;default:1;comment:文件类型 1文件 2图片 3文档" json:"type" form:"type"`
	Size      int64     `gorm:"column:size;type:bigint;comment:文件大小" json:"size" form:"size"`
	Name      string    `gorm:"column:name;type:varchar(150);comment:原来文件名" json:"name" form:"name"`
	Sha256    string    `gorm:"column:sha256;type:varchar(70);index;comment:sha256值" json:"sha256" form:"sha256"`
	Path      string    `gorm:"column:path;type:varchar(150);comment:文件路径" json:"path" form:"path"`
	Mime      string    `gorm:"column:mime;type:varchar(30);comment:文件Mime类型" json:"mime" form:"mime"`
	Suffix    string    `gorm:"column:suffix;type:varchar(20);comment:文件后缀" json:"suffix" form:"suffix"`
	Uid       int64     `gorm:"column:uid;type:bigint;comment:上传用户" json:"uid" form:"uid"`
	Private   int8      `gorm:"column:private;type:tinyint;default:0;not null;comment:是否为私有" json:"private" form:"private"`
	Media     int8      `gorm:"column:media;type:TINYINT;;default:0;comment:存储介质" json:"media" form:"media"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"created_at" form:"created_at"`
}
