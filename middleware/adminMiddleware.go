package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/models"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exist := c.Get("user_type")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User type not found"})
			c.Abort()
			return
		}

		// Check if the user is an admin
		if userType != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this route"})
			c.Abort()
			return
		}

		userId, exist := c.Get("user_id")

		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User id not found"})
			c.Abort()
			return
		}
		// Check if the admin exists
		var admin = models.Admin{
			UserID: userId.(int64),
		}
		if err := admin.GetAdminByUserID(); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Admin not found"})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}
