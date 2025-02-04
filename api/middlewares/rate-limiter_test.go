package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(RateLimiter(2)) // 2 requêtes par seconde
	router.GET("/limited", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Envoi de 2 requêtes valides
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/limited", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	}

	// Envoi d'une 3e requête (doit être bloquée)
	req := httptest.NewRequest("GET", "/limited", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.Contains(t, w.Body.String(), "too many requests")

	// Attente pour ne plus être bloqué
	time.Sleep(time.Second)

	// Nouvelle requête après attente (doit passer)
	req = httptest.NewRequest("GET", "/limited", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
