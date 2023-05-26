package main

import (
	"fmt"
	"github.com/MrRytis/go-fiber-test/migration"
	"github.com/MrRytis/go-fiber-test/model"
	"github.com/MrRytis/go-fiber-test/utils"
	"log"
)

func main() {
	db := utils.NewDb()

	count := 0

	migrations := migration.GetAllMigrations()
	for _, m := range migrations {
		if migration.IsMigrated(m.Name(), db) {
			err := m.Down(db)
			if err != nil {
				log.Fatal(err)
			}

			if err := db.Where("name = ?", m.Name()).Delete(&model.Migration{}).Error; err != nil {
				log.Fatal(err)
			}

			count++

			fmt.Println("Rollback: " + m.Name())
		}
	}

	fmt.Println("Rollback " + fmt.Sprint(count) + " migrations")
	fmt.Println("Rollbacks completed")
}
