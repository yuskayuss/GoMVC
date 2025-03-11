package routes

import (
	"golang-mvc-postgresql/controllers" // Seharusnya ini, bukan controllers_test

	"github.com/gin-gonic/gin"
)

// Setup router untuk test & aplikasi utama
func SetupRouter() *gin.Engine {
	router := gin.Default()
	RegisterRoutes(router)
	return router
}

// Fungsi untuk register semua routes
func RegisterRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.POST("/register", controllers.RegisterUser) // Tambahkan endpoint register
}
