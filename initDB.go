package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库连接
func initDB() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPassword, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Movie{})
	if err != nil {
		return nil, err
	} // 创建电影表
	//fmt.Println("表创建成功")

	return db, err
}
