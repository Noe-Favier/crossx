package auth

import (
	"crossx/models"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// Test de HashPassword et CheckPasswordHash
func TestPasswordHashing(t *testing.T) {
	password := "mypassword"

	// Hash du mot de passe
	hash, err := HashPassword(password)
	assert.NoError(t, err, "HashPassword ne doit pas renvoyer d'erreur")
	assert.NotEmpty(t, hash, "Le hash ne doit pas être vide")

	// Vérification du hash
	match := CheckPasswordHash(password, hash)
	assert.True(t, match, "Le mot de passe doit correspondre au hash")

	// Test avec un mauvais mot de passe
	wrongMatch := CheckPasswordHash("wrongpassword", hash)
	assert.False(t, wrongMatch, "Un mauvais mot de passe ne doit pas correspondre")
}

// Test de GenerateJWT
func TestGenerateJWT(t *testing.T) {
	// Set fake JWT secret
	os.Setenv("JWT_SECRET", "mysecret")

	user := models.User{ID: 1, Email: "test@example.com"}
	token, err := GenerateJWT(user)

	assert.NoError(t, err, "GenerateJWT ne doit pas renvoyer d'erreur")
	assert.NotEmpty(t, token, "Le token JWT ne doit pas être vide")

	// Vérification du format du token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	assert.NoError(t, err, "Le token généré doit être valide")
	assert.True(t, parsedToken.Valid, "Le token doit être valide")
}

// Test de ValidateToken
func TestValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "mysecret")

	user := models.User{ID: 1, Email: "test@example.com"}
	token, _ := GenerateJWT(user)

	parsedToken, err := ValidateToken(token)
	assert.NoError(t, err, "ValidateToken ne doit pas renvoyer d'erreur")
	assert.NotNil(t, parsedToken, "Le token ne doit pas être nil")
	assert.True(t, parsedToken.Valid, "Le token doit être valide")
}

// Test d'un token invalide
func TestValidateInvalidToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "mysecret")

	_, err := ValidateToken("invalid.token.here")
	assert.Error(t, err, "Un token invalide doit renvoyer une erreur")
}
