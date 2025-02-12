package auth

import (
	"errors"
	"os"
	"time"

	"crossx/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash un mot de passe
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash vérifie la correspondance mot de passe/hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT génère un token JWT
func GenerateJWT(user models.User) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET")) // Définissez cette variable d'environnement

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(secret)
}

// ValidateTokenFunc est la fonction utilisée pour valider le token.
// Par défaut, elle pointe sur la fonction ValidateToken, mais elle pourra être remplacée en test.
var ValidateTokenFunc = ValidateToken

// HashPassword, CheckPasswordHash, GenerateJWT, ValidateToken restent inchangés...

// ValidateToken vérifie la validité d'un token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return secret, nil
	})
}

func PartialUserFromToken(token *jwt.Token) models.User {
	tmp := token.Claims.(jwt.MapClaims)
	return models.User{Username: tmp["username"].(string), ID: uint(tmp["id"].(float64)), Email: tmp["email"].(string)}
}
