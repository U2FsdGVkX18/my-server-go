package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	logger "my-server-go/tool/log"
	"time"
)

var (
	Client *redis.Client
	ctx    = context.Background()
)

func init() {
	Connect()
}

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "125.91.35.185:6379",
		Password: "mujin1110",
		DB:       0,
	})
}

func GetValue(key string) string {
	result, err := Client.Get(ctx, key).Result()
	if err != nil {
		logger.Write("GetValue redis获取值错误或为空值:", err)
		return ""
	}
	return result
}

func SetValue(key string, value any, expiration time.Duration) {
	err := Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		logger.Write("SetValue redis设置值错误:", err)
	}
}
