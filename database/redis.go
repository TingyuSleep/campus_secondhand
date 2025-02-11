package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Rdb *redis.Client
var Ctx = context.Background()

// InitRedis 初始化 Redis 连接
func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // Redis 密码
		DB:       0,                // 使用默认数据库
	})

	// 测试 Redis 连接
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Redis 连接失败: %v", err)
	}
	log.Println("Redis 连接成功")
}
