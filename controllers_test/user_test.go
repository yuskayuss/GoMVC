package controllers_test

import (
	"golang-mvc-postgresql/config"
	"golang-mvc-postgresql/controllers"
	"golang-mvc-postgresql/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup database untuk testing
// func setupTestDatabase() {
// 	config.ConnectDatabase() 
// 	config.DB.Exec("DROP TABLE IF EXISTS users") 
// 	config.DB.AutoMigrate(&models.User{})       
// }

func setupTestDatabase() {
	config.ConnectDatabase() // Pastikan database dikoneksikan sebelum test
	config.DB.Exec("DROP TABLE IF EXISTS users") // Hapus tabel users jika sudah ada
	config.DB.AutoMigrate(&models.User{})       // Buat ulang tabel users

	// Tambahkan data dummy
	users := []models.User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
	}

	for _, user := range users {
		config.DB.Create(&user)
	}
}


// Buat router tanpa mengimpor routes
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	return router
}

func TestGetUsers(t *testing.T) {
	setupTestDatabase()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Gunakan StatusCreated (201) jika user berhasil dibuat
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

// RegisterUser menangani registrasi user baru
func RegisterUser(c *gin.Context) {
	var user models.User

	// Bind JSON ke struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendaftarkan user"})
		return
	}

	// Response sukses
	c.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil", "user": user})
}

