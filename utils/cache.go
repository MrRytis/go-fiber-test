package utils

import (
	"github.com/redis/go-redis/v9"
	"os"
)

func NewCache() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
