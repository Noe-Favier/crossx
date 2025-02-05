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

// setupTestDBPost initialise une base SQLite en mémoire pour les tests
func setupTestDBPost() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Post{}) // Migration de la table Post
	database.SetTestDB(db)         // Injection de la base dans database
}

// setupTestRouterPost configure le routeur Gin pour les tests
func setupTestRouterPost() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/post/:id", GetPost)
	r.POST("/api/v1/post", CreatePost)
	r.PUT("/api/v1/post/:id", UpdatePost)
	r.DELETE("/api/v1/post/:id", DeletePost)
	return r
}

// TestGetPost_NotFound vérifie qu'un post inexistant retourne 404
func TestGetPost_NotFound(t *testing.T) {
	setupTestDBPost()

	r := setupTestRouterPost()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/post/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestCreatePost vérifie la création d'un post
func TestCreatePost(t *testing.T) {
	setupTestDBPost()

	r := setupTestRouterPost()
	w := httptest.NewRecorder()

	post := models.Post{Content: "Test Post", UserID: 1}
	jsonValue, _ := json.Marshal(post)

	req, _ := http.NewRequest("POST", "/api/v1/post", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestUpdatePost_NotFound vérifie qu'on ne peut pas mettre à jour un post inexistant
func TestUpdatePost_NotFound(t *testing.T) {
	setupTestDBPost()

	r := setupTestRouterPost()
	w := httptest.NewRecorder()

	post := models.Post{Content: "Updated Content"}
	jsonValue, _ := json.Marshal(post)

	req, _ := http.NewRequest("PUT", "/api/v1/post/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestDeletePost_NotFound vérifie qu'on ne peut pas supprimer un post inexistant
func TestDeletePost_NotFound(t *testing.T) {
	setupTestDBPost()

	r := setupTestRouterPost()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/api/v1/post/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
