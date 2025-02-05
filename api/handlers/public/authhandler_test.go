package public

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"crossx/auth"
	"crossx/database"
	"crossx/models"
	"crossx/models/dto"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDBPublic initialise une base SQLite en mémoire et migre le modèle User.
func setupTestDBPublic() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	database.SetTestDB(db)
}

// setupTestRouterPublicAuth configure un routeur Gin pour les endpoints d'authentification.
func setupTestRouterPublicAuth() *gin.Engine {
	r := gin.Default()
	// Définition des routes pour Login et Signup
	r.POST("/api/v1/public/login", Login)
	r.POST("/api/v1/public/signup", Signup)
	return r
}

// TestSignupEndpoint vérifie que l'inscription fonctionne correctement et retourne un token et les infos de l'utilisateur.
func TestSignupEndpoint(t *testing.T) {
	setupTestDBPublic()
	router := setupTestRouterPublicAuth()

	// Prépare un payload conforme au dto.SignupForm attendu par l'endpoint Signup.
	payload := dto.SignupForm{
		Username: "TestUser",
		Email:    "test@example.com",
		Password: "testpassword",
	}
	jsonValue, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/public/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Vérifie que l'inscription retourne bien un code 200.
	assert.Equal(t, http.StatusOK, w.Code, "L'inscription doit réussir")

	// Vérifie que la réponse est un JSON contenant un token et les informations de l'utilisateur.
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "La réponse doit être un JSON valide")
	assert.NotNil(t, response["token"], "Un token JWT doit être généré")
	assert.NotNil(t, response["user"], "Les informations de l'utilisateur doivent être retournées")
}

// TestLoginEndpoint vérifie que la connexion fonctionne correctement avec un utilisateur existant.
func TestLoginEndpoint(t *testing.T) {
	setupTestDBPublic()
	router := setupTestRouterPublicAuth()

	// Crée un utilisateur de test avec un mot de passe haché.
	hashedPassword, _ := auth.HashPassword("testpassword")
	user := models.User{
		Username:     "ExistingUser",
		Email:        "existing@example.com",
		PasswordHash: hashedPassword,
	}
	db := database.GetDB()
	db.Create(&user)

	// Prépare un payload conforme au dto.LoginForm attendu par l'endpoint Login.
	payload := dto.LoginForm{
		Username: "ExistingUser",
		Password: "testpassword",
	}
	jsonValue, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/public/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Vérifie que la connexion retourne bien un code 200.
	assert.Equal(t, http.StatusOK, w.Code, "La connexion doit réussir")

	// Vérifie que la réponse contient un token et les informations de l'utilisateur.
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "La réponse doit être un JSON valide")
	assert.NotNil(t, response["token"], "Un token JWT doit être généré")
	assert.NotNil(t, response["user"], "Les informations de l'utilisateur doivent être retournées")
}
