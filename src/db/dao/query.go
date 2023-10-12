package dao

import (
	"basego/src/db"
	"database/sql"

	"gorm.io/gorm"
)

type QueryCondition struct {
	Column string
	Input  interface{}
}

func QueryObjectList(gormDb *gorm.DB, object db.DbModel, page, pageSize int32,
	qc ...*QueryCondition) (sqlRow *sql.Rows, err error) {

	offset := (page - 1) * pageSize

	querySub := gormDb.Model(object).Select("id").
		Limit(int(pageSize)).Offset(int(offset)).Order("id desc")

	if qc != nil {
		for i := 0; i < len(qc); i++ {
			if qc[i] != nil {
				querySub = querySub.Where(qc[i].Column+" = ?", qc[i].Input)
			}
		}
	}

	return gormDb.Model(object).Order("id desc").Joins("INNER JOIN (?) AS t2 USING (id)", querySub).Rows()
}

func QueryObjectListTotal(gormDb *gorm.DB, object db.DbModel,
	totalChan chan int64, qc ...*QueryCondition) {

	var total int64

	totalSub := gormDb.Model(object)

	if qc != nil {
		for i := 0; i < len(qc); i++ {
			if qc[i] != nil {
				totalSub = totalSub.Where(qc[i].Column+" = ?", qc[i].Input)
			}
		}
	}

	totalSub.Count(&total)

	totalChan <- total

}

func QueryObjectByCondition(gormDb *gorm.DB, object db.DbModel,
	conditions ...*QueryCondition) error {

	db := gormDb.Model(object)

	for _, c := range conditions {
		db = db.Where(c.Column+" = ?", c.Input)
	}

	if err := db.First(object).Error; err != nil {
		return err
	}

	return nil
}
