package main

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

var (
	rdb         *redis.Client
	db          *gorm.DB
	bloomFilter *bloom.BloomFilter
	bloomMutex  sync.Mutex
	bloommutex  sync.Mutex
)
