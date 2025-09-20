package models

// Option 选项表
type Option struct {
	Id     int32  `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:业务ID" json:"id" form:"id"`
	Key    string `gorm:"column:key;type:varchar(50);unique;not null;comment:选项Key" json:"key" form:"key"`
	Value  string `gorm:"column:value;type:text;comment:选项值" json:"value" form:"value"`
	Remark string `gorm:"column:remark;type:varchar(50);comment:备注" json:"remark" form:"remark"`
}
