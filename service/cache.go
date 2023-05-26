package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func SetCache(cache *redis.Client, key string, values interface{}, duration time.Duration) {
	cache.Set(context.Background(), key, values, duration)
}

func GetCache(cache *redis.Client, key string) {

}
