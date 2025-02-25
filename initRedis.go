package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func initRedis() {
	// 初始化Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Println("Redis连接失败:", err)

	}
	//defer rdb.Close()
}
