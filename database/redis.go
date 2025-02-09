package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Redis连接失败：%v", err)
	}
	log.Fatalf("Redis连接成功")
}
