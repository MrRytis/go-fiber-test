package migration

import (
	"github.com/MrRytis/go-fiber-test/model"
	"gorm.io/gorm"
)

type Migration20230518192600 struct {
}

func (m *Migration20230518192600) Up(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{})
}

func (m *Migration20230518192600) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&model.User{})
}

func (m *Migration20230518192600) Name() string {
	return "Migration_20230518192600"
}
