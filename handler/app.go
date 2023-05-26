package handler

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewApp(db *gorm.DB, cache *redis.Client) *App {
	return &App{
		db:    db,
		cache: cache,
	}
}
