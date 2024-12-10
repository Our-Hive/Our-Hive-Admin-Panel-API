package security

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoleMiddleware(c *gin.Context) {
	user, ok := c.Get("user")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return

	}

	if user.(model.Token).Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Forbidden, you are not an admin",
		})
		return
	}
	c.Next()
}
