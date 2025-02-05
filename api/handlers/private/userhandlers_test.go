package private

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"crossx/database"
	"crossx/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDBUser initialise une base SQLite en mémoire pour les tests
func setupTestDBUser() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{}) // Migration de la table User
	database.SetTestDB(db)         // Injection de la base dans database
}

// setupTestRouterUser configure le routeur Gin pour les tests
func setupTestRouterUser() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/user/:id", GetUser)
	r.POST("/api/v1/user", CreateUser)
	r.PUT("/api/v1/user/:id", UpdateUser)
	r.DELETE("/api/v1/user/:id", DeleteUser)
	return r
}

// TestGetUser_NotFound vérifie qu'un utilisateur inexistant retourne 404
func TestGetUser_NotFound(t *testing.T) {
	setupTestDBUser()

	r := setupTestRouterUser()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestCreateUser vérifie la création d'un utilisateur
func TestCreateUser(t *testing.T) {
	setupTestDBUser()

	r := setupTestRouterUser()
	w := httptest.NewRecorder()

	user := models.User{Username: "TestUser", Email: "test@example.com", PasswordHash: "hashedpassword"}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestUpdateUser_NotFound vérifie qu'on ne peut pas mettre à jour un utilisateur inexistant
func TestUpdateUser_NotFound(t *testing.T) {
	setupTestDBUser()

	r := setupTestRouterUser()
	w := httptest.NewRecorder()

	user := models.User{Username: "UpdatedUser", Email: "updated@example.com"}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("PUT", "/api/v1/user/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestDeleteUser_NotFound vérifie qu'on ne peut pas supprimer un utilisateur inexistant
func TestDeleteUser_NotFound(t *testing.T) {
	setupTestDBUser()

	r := setupTestRouterUser()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/api/v1/user/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
