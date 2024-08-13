// initializers/database.go

package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DB_URL")
	var err error
	if db == "postgres" {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
}
