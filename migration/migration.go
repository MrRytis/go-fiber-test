package migration

import (
	"github.com/MrRytis/go-fiber-test/model"
	"gorm.io/gorm"
	"log"
)

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
	Name() string
}

func GetAllMigrations() []Migration {
	return []Migration{
		&Migration20230518192600{},
	}
}

func IsMigrated(name string, db *gorm.DB) bool {
	var m model.Migration
	if err := db.Where("name = ?", name).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}

		log.Fatal(err.Error())
	}

	return false
}
