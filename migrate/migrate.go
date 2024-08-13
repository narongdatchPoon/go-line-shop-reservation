// migrate/migrate.go

package main

import (
	"go-line-shop-reservation/initializers"
	"go-line-shop-reservation/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Reservation{})
}
