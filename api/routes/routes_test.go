package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test de l'initialisation du routeur
func TestNew(t *testing.T) {
	r := New()
	assert.NotNil(t, r, "New() ne doit pas retourner nil")
	assert.NotNil(t, r.router, "Le routeur doit être initialisé")
}

// Test du routeur avec des routes publiques
func TestSetupRouter_PublicRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode) // Mode test pour éviter les logs inutiles
	r := New()
	router := r.SetupRouter()

	// Vérifier la route /api/v1/public/health
	req, _ := http.NewRequest("GET", "/api/v1/public/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "La route /health doit retourner un statut 200")

	// Vérifier la route Swagger
	req, _ = http.NewRequest("GET", "/swagger/index.html", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "La route /swagger doit retourner un statut 200")
}

// Test d'une route protégée (requiert un middleware d'authentification)
func TestSetupRouter_ProtectedRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := New()
	router := r.SetupRouter()

	// Simuler un token JWT valide pour bypasser l'auth middleware
	req, _ := http.NewRequest("GET", "/api/v1/test", nil)
	req.Header.Set("Authorization", "Bearer test_token") // Simule un token

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code, "Une route protégée sans authentification valide doit être refusée")
}
