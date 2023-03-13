package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	logger "my-server-go/tool/log"
	"time"
)

var ctx = context.Background()

func Connect() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "125.91.35.185:6379",
		Password: "mujin1110",
		DB:       0,
	})
	return rdb
}

func GetValue(key string) string {
	rdb := Connect()
	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		logger.Write("GetValue redis获取值错误:", err)
	}
	return result
}

func SetValue(key string, value any, expiration time.Duration) {
	rdb := Connect()
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		logger.Write("SetValue redis设置值错误:", err)
	}
}
