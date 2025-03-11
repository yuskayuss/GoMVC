package controllers

import (
	"golang-mvc-postgresql/config"
	"golang-mvc-postgresql/routes"

	"github.com/gin-gonic/gin"
)

// Inisialisasi database sebelum semua test
func init() {
	config.ConnectDatabase()
}

// Setup router untuk testing
func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.RegisterRoutes(router) // Panggil RegisterRoutes dari package routes
	return router
}
