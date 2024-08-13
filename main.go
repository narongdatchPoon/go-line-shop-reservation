// main.go

package main

import (
	"go-line-shop-reservation/controllers"
	"go-line-shop-reservation/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/line", controllers.WebHook)

	router.Run()
}
