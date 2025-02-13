package private

import (
	"crossx/database"
	"crossx/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Consultation d'un commentaire
// @Description	Récupère un commentaire par son ID
// @Tags			comments
// @Produce		json
// @Param			id		path		int	true	"ID du commentaire"
// @Success		200		{object}	models.Comment
// @Failure		404		{object}	map[string]string	"Comment not found"
// @Failure		500		{object}	map[string]string	"Erreur interne"
// @Router			/comments/{id} [get]
// @Security		ApiKeyAuth
func GetComment(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var comment models.Comment

	if err := db.Preload("Post").Preload("User").First(&comment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, comment)
}

// @Summary		Création d'un commentaire
// @Description	Ajoute un nouveau commentaire
// @Tags			comments
// @Accept			json
// @Produce		json
// @Param			body	body		models.Comment	true	"Corps du commentaire"
// @Success		201		{object}	models.Comment
// @Failure		400		{object}	map[string]string	"Invalid input"
// @Failure		500		{object}	map[string]string	"Erreur interne"
// @Router			/comments [post]
// @Security		ApiKeyAuth
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	user := c.MustGet("user").(models.User)
	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = user.ID

	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// @Summary		Modification d'un commentaire
// @Description	Met à jour un commentaire existant
// @Tags			comments
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"ID du commentaire"
// @Param			body	body		models.Comment	true	"Corps du commentaire"
// @Success		200		{object}	models.Comment
// @Failure		400		{object}	map[string]string	"Invalid input"
// @Failure		404		{object}	map[string]string	"Comment not found"
// @Failure		500		{object}	map[string]string	"Erreur interne"
// @Router			/comments/{id} [put]
// @Security		ApiKeyAuth
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var comment models.Comment

	if err := db.First(&comment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Content = input.Content
	comment.PostID = input.PostID
	comment.UserID = input.UserID

	if err := db.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// @Summary		Suppression d'un commentaire
// @Description	Supprime un commentaire par son ID
// @Tags			comments
// @Produce		json
// @Param			id		path		int	true	"ID du commentaire"
// @Success		200		{object}	map[string]string	"Comment deleted successfully"
// @Failure		404		{object}	map[string]string	"Comment not found"
// @Failure		500		{object}	map[string]string	"Erreur interne"
// @Router			/comments/{id} [delete]
// @Security		ApiKeyAuth
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var comment models.Comment

	if err := db.First(&comment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
