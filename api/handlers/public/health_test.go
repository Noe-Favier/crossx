package public

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupTestRouterPublicHealth configure un routeur Gin pour l'endpoint de health.
func setupTestRouterPublicHealth() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/public/health", HealthHandler)
	return r
}

// TestHealthEndpoint vérifie que l'endpoint de health check retourne bien un statut 200 et le message attendu.
func TestHealthEndpoint(t *testing.T) {
	router := setupTestRouterPublicHealth()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/public/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "La réponse doit être un JSON valide")
	assert.Equal(t, "ok", response["status"], "Le status doit être 'ok'")
}
