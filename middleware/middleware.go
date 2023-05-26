package middleware

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Middleware struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewMiddleware(db *gorm.DB, cache *redis.Client) *Middleware {
	return &Middleware{
		db:    db,
		cache: cache,
	}
}
