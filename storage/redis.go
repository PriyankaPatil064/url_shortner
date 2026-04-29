package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis connected ✅")
}
