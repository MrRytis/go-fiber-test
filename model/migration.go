package model

import "time"

type Migration struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	ExecutedAt time.Time
}
