package private

import (
	"crossx/database"
	"crossx/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// GetUser - Consultation d'un utilisateur
//
//	@Summary		Consultation d'un utilisateur
//	@Description	Permet de récupérer les détails d'un utilisateur via son ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"ID de l'utilisateur"
//	@Success		200	{object}	models.User
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/users/{id} [get]
//	@Security		ApiKeyAuth
func GetUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser - Création d'un utilisateur
//
//	@Summary		Création d'un utilisateur
//	@Description	Permet de créer un nouvel utilisateur
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true	"Informations de l'utilisateur"
//	@Success		201	{object}	models.User
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/users [post]
//	@Security		ApiKeyAuth
func CreateUser(c *gin.Context) {
	db := database.GetDB()

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// UpdateUser - Modification d'un utilisateur
//
//	@Summary		Modification d'un utilisateur
//	@Description	Permet de modifier les informations d'un utilisateur existant
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"ID de l'utilisateur"
//	@Param			user	body		models.User	true	"Nouvelles informations de l'utilisateur"
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/users/{id} [put]
//	@Security		ApiKeyAuth
func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Bio = input.Bio
	user.Email = input.Email
	user.Username = input.Username
	user.PasswordHash = input.PasswordHash
	user.ProfilePictureUrl = input.ProfilePictureUrl

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser - Suppression d'un utilisateur
//
//	@Summary		Suppression d'un utilisateur
//	@Description	Permet de supprimer un utilisateur via son ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"ID de l'utilisateur"
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/users/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
