package db

type DbModel interface {
	TableName() string
}

var TableSlice = make([]interface{}, 0)

type GeneralField struct {
	Id         int32 `gorm:"primaryKey;autoIncrement" json:"id"`
	CreateTime int64 `gorm:"autoCreateTime:milli" json:"createTime"`
	UpdateTime int64 `gorm:"autoUpdateTime:milli" json:"updateTime"`
}
