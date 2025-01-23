package private

import (
	"crossx/database"
	"crossx/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// GetPost - Consultation d'un post
//
//	@Summary		Consultation d'un post
//	@Description	Récupère un post par son ID
//	@Tags			posts
//	@Produce		json
//	@Param			id		path		int	true	"ID du post"
//	@Success		200		{object}	models.Post
//	@Failure		404		{object}	map[string]string	"Post not found"
//	@Failure		500		{object}	map[string]string	"Erreur interne"
//	@Router			/posts/{id} [get]
//	@Security		ApiKeyAuth
func GetPost(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var post models.Post

	if err := db.Preload("User").First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, post)
}

// CreatePost - Création d'un post
//
//	@Summary		Création d'un post
//	@Description	Ajoute un nouveau post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.Post	true	"Corps du post"
//	@Success		201		{object}	models.Post
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		500		{object}	map[string]string	"Erreur interne"
//	@Router			/posts [post]
//	@Security		ApiKeyAuth
func CreatePost(c *gin.Context) {
	db := database.GetDB()
	var input models.Post
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

// UpdatePost - Modification d'un post
//
//	@Summary		Modification d'un post
//	@Description	Met à jour un post existant
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"ID du post"
//	@Param			body	body		models.Post	true	"Corps du post"
//	@Success		200		{object}	models.Post
//	@Failure		400		{object}	map[string]string	"Invalid input"
//	@Failure		404		{object}	map[string]string	"Post not found"
//	@Failure		500		{object}	map[string]string	"Erreur interne"
//	@Router			/posts/{id} [put]
//	@Security		ApiKeyAuth
func UpdatePost(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var post models.Post

	if err := db.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Content = input.Content
	post.MediaUrl = input.MediaUrl
	post.UserID = input.UserID

	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost - Suppression d'un post
//
//	@Summary		Suppression d'un post
//	@Description	Supprime un post par son ID
//	@Tags			posts
//	@Produce		json
//	@Param			id		path		int	true	"ID du post"
//	@Success		200		{object}	map[string]string	"Post deleted successfully"
//	@Failure		404		{object}	map[string]string	"Post not found"
//	@Failure		500		{object}	map[string]string	"Erreur interne"
//	@Router			/posts/{id} [delete]
//	@Security		ApiKeyAuth
func DeletePost(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	var post models.Post

	if err := db.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
