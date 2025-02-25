package main

import (
	"time"

	"gorm.io/gorm"
)

// 电影数据结构
type Movie struct {
	gorm.Model
	Title     string `gorm:"unique,not null"`
	Year      string //`gorm:"unique"`
	Rating    string
	Directors []string `gorm:"type:json;serializer:json"`
	Actors    []string `gorm:"type:json;serializer:json"`
	Genre     []string `gorm:"type:json;serializer:json"`
}

// Redis连接配置
const (
	redisAddr = "localhost:6379"
	//bloomFilter = "movie_filter"
	MaxDepth    = 3
	parallelism = 2
	RandomDelay = 5 * time.Second
)

// MySQL连接配置
const (
	dbUser     = "root"
	dbPassword = "3120004654"
	dbName     = "movie_db"
)
