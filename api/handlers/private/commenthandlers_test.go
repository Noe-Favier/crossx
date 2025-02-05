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

// setupTestDBComment initialise une base SQLite en mémoire pour les tests
// et l'assigne au package database via SetTestDB.
func setupTestDBComment() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Erreur lors de l'ouverture de la base de données : " + err.Error())
	}
	// Migration de la table Comment
	if err := db.AutoMigrate(&models.Comment{}); err != nil {
		panic("Erreur lors de l'auto-migration : " + err.Error())
	}
	database.SetTestDB(db)
}

// setupTestRouterComment configure le routeur Gin pour les tests.
func setupTestRouterComment() *gin.Engine {
	// Passe en mode release pour éviter les logs de debug pendant les tests.
	gin.SetMode(gin.ReleaseMode)
	r := gin.New() // On utilise gin.New() pour éviter l'ajout automatique de middlewares.
	r.GET("/api/v1/comment/:id", GetComment)
	r.POST("/api/v1/comment", CreateComment)
	r.PUT("/api/v1/comment/:id", UpdateComment)
	r.DELETE("/api/v1/comment/:id", DeleteComment)
	return r
}

// TestGetComment_NotFound vérifie qu'un commentaire inexistant retourne 404.
func TestGetComment_NotFound(t *testing.T) {
	setupTestDBComment()
	router := setupTestRouterComment()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/comment/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestCreateComment vérifie la création d'un commentaire.
func TestCreateComment(t *testing.T) {
	setupTestDBComment()
	router := setupTestRouterComment()
	w := httptest.NewRecorder()

	comment := models.Comment{Content: "Test Comment", PostID: 1, UserID: 1}
	jsonValue, err := json.Marshal(comment)
	if err != nil {
		t.Fatalf("Erreur lors du marshalling JSON: %v", err)
	}

	req, _ := http.NewRequest("POST", "/api/v1/comment", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestUpdateComment_NotFound vérifie qu'on ne peut pas mettre à jour un commentaire inexistant.
func TestUpdateComment_NotFound(t *testing.T) {
	setupTestDBComment()
	router := setupTestRouterComment()
	w := httptest.NewRecorder()

	comment := models.Comment{Content: "Updated Content"}
	jsonValue, err := json.Marshal(comment)
	if err != nil {
		t.Fatalf("Erreur lors du marshalling JSON: %v", err)
	}

	req, _ := http.NewRequest("PUT", "/api/v1/comment/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestDeleteComment_NotFound vérifie qu'on ne peut pas supprimer un commentaire inexistant.
func TestDeleteComment_NotFound(t *testing.T) {
	setupTestDBComment()
	router := setupTestRouterComment()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/comment/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
