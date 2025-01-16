package private

import "github.com/gin-gonic/gin"

//	@Summary		Test protected endpoint
//	@Description	Test endpoint requiring authentication
//	@Tags			protected
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/test [get]
//	@Security		ApiKeyAuth
func TestHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "protected test"})
}
