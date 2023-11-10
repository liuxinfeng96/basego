package db

type DbModel interface {
	TableName() string
}

// TableSlice gorm自动建表模型列表
var TableSlice = make([]interface{}, 0)

// GeneralField 表结构通用字段
type GeneralField struct {
	Id         int32 `gorm:"primaryKey;autoIncrement" json:"id"`
	CreateTime int64 `gorm:"autoCreateTime:milli" json:"createTime"`
	UpdateTime int64 `gorm:"autoUpdateTime:milli" json:"updateTime"`
}
