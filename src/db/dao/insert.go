package dao

import (
	"basego/src/db"

	"gorm.io/gorm"
)

func InsertObjectsToDBInTransaction(gormDb *gorm.DB, objects []db.DbModel) error {
	tx := gormDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for i := range objects {
		if err := tx.Create(objects[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func InsertOneObjectToDB(object db.DbModel, gormDb *gorm.DB) error {
	if err := gormDb.Create(object).Error; err != nil {
		return err
	}
	return nil
}
