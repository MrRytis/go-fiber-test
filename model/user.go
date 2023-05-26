package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uid           string `gorm:"uniqueIndex"`
	Email         string `gorm:"uniqueIndex"`
	Password      string
	Name          string
	Surname       string
	Icon          *string
	EmailVerified bool `gorm:"default:false"`
	Reset         bool `gorm:"default:false"`
	GoogleId      *string
	FacebookId    *string
}
