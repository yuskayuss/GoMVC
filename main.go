package main

import (
	"fmt"
	"golang-mvc-postgresql/config"
	"golang-mvc-postgresql/models"
	"golang-mvc-postgresql/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Auto-migrate tabel
	config.DB.AutoMigrate(&models.User{})

	router := gin.Default()
	routes.RegisterRoutes(router)

	fmt.Println("Server running on port 8080")
	router.Run(":9090")
}
