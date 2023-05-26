package main

import (
	"fmt"
	"github.com/MrRytis/go-fiber-test/migration"
	"github.com/MrRytis/go-fiber-test/model"
	"github.com/MrRytis/go-fiber-test/utils"
	"log"
	"time"
)

func main() {
	db := utils.NewDb()

	db.AutoMigrate(&model.Migration{})

	count := 0

	migrations := migration.GetAllMigrations()
	for _, m := range migrations {
		if !migration.IsMigrated(m.Name(), db) {
			err := m.Up(db)
			if err != nil {
				log.Fatal(err)
			}

			mig := model.Migration{
				Name:       m.Name(),
				ExecutedAt: time.Now(),
			}

			if err := db.Create(&mig).Error; err != nil {
				log.Fatal(err)
			}

			count++

			fmt.Println("Migrated: " + m.Name())
		}
	}

	fmt.Println("Migrated " + fmt.Sprint(count) + " migrations")
	fmt.Println("Migrations completed")
}
