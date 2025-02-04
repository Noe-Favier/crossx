package public

import (
	"bytes"
	"crossx/models/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", Login)

	loginForm := dto.LoginForm{
		Username: "testuser",
		Password: "testpass",
	}

	jsonValue, _ := json.Marshal(loginForm)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code, "L'utilisateur inexistant doit retourner unauthorized")

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Invalid credentials", response["error"], "Le message d'erreur doit être 'Invalid credentials'")
}

func TestSignupEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/signup", Signup)
	os.Setenv("JWT_SECRET", "test_secret")

	signupForm := dto.SignupForm{
		Username: "newuser",
		Email:    "test@test.com",
		Password: "testpass",
	}

	jsonValue, _ := json.Marshal(signupForm)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "L'inscription doit réussir")

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"], "Un token JWT doit être généré")
	assert.NotNil(t, response["user"], "Les informations de l'utilisateur doivent être retournées")
}
