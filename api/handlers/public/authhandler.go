package public

import (
	"crossx/auth"
	"crossx/database"
	"crossx/models"
	"crossx/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDB()
	var input dto.LoginForm

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !auth.CheckPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func Signup(c *gin.Context) {
	db := database.GetDB()
	var input dto.SignupForm

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	hash, err := auth.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	user := models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: hash,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
