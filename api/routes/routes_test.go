package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestNew vérifie simplement que New() retourne bien une instance de Routes et que le routeur est initialisé.
func TestNew(t *testing.T) {
	r := New()
	assert.NotNil(t, r, "New() ne doit pas retourner nil")
	assert.NotNil(t, r.router, "Le routeur doit être initialisé")
}

// TestSetupRouter_PublicRoutes vérifie que les routes publiques retournent bien les statuts attendus.
func TestSetupRouter_PublicRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode) // Mode test pour éviter les logs inutiles
	r := New()
	router := r.SetupRouter()

	// Test de la route /api/v1/public/health
	req, _ := http.NewRequest("GET", "/api/v1/public/health", nil)
	req.Host = "localhost:8080" // Pour satisfaire le middleware secure
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "La route /health doit retourner un statut 200")

	// Vérifier le contenu de la réponse
	var response map[string]string
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, "ok", response["status"], "Le status doit être 'ok'")
}

// TestSetupRouter_ProtectedRoutes vérifie qu'une route protégée sans authentification retourne bien 401.
func TestSetupRouter_ProtectedRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := New()
	router := r.SetupRouter()

	// Test des routes protégées sans token
	protectedRoutes := []string{
		"/api/v1/comment/1",
		"/api/v1/post/1",
		"/api/v1/user/1",
	}

	for _, route := range protectedRoutes {
		req, _ := http.NewRequest("GET", route, nil)
		req.Host = "localhost:8080"
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code,
			"La route %s sans authentification doit être refusée", route)
	}
}
