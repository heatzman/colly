package main

import (
	"errors"
	"gorm.io/gorm/clause"
)

// 保存电影数据到MySQL
func saveMovie(m Movie) error {

	if db == nil {
		return errors.New("db not initialized")
	}
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "title"}, {Name: "year"}}, // 冲突检测字段
		DoNothing: true,                                             // 冲突时忽略
	}).Create(&m)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
