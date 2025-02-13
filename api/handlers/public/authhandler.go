package public

import (
	"crossx/auth"
	"crossx/database"
	"crossx/models"
	"crossx/models/dto"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
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

	// Parse multipart form
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10 MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	// Get form values
	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	// Handle profile image
	file, header, err := c.Request.FormFile("profileImage")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process profile image"})
		return
	}

	var profilePictureUrl string
	if file != nil {
		defer file.Close()

		// Create uploads directory if it doesn't exist
		uploadDir := "uploads/profiles"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		// Generate unique filename
		ext := filepath.Ext(header.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath := filepath.Join(uploadDir, filename)

		// Save the file
		if err := c.SaveUploadedFile(header, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		// Convert file path to URL
		profilePictureUrl = fmt.Sprintf("/api/v1/uploads/profiles/%s", filename)
	}

	// Hash password
	hash, err := auth.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	// Create user
	user := models.User{
		Username:          username,
		Email:             email,
		PasswordHash:      hash,
		ProfilePictureUrl: profilePictureUrl,
	}

	if err := db.Create(&user).Error; err != nil {
		// Clean up uploaded file if user creation fails
		if profilePictureUrl != "" {
			filePath := filepath.Join("uploads/profiles", filepath.Base(profilePictureUrl))
			os.Remove(filePath)
		}

		if strings.Contains(err.(*pgconn.PgError).Message, "duplicate key value violates unique constraint") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already taken"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}

	// Generate JWT
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

func Me(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userModel, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, userModel)
}
