package middleware

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		//  gett user from access token
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not found"})
			c.Abort()
			return
		}

		// Get the user type from the context
		userTypeCtx, exist := c.Get("user_type")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User type not found"})
			c.Abort()
			return
		}

		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			c.Abort()
			return
		}

		//
		userType := claims.User_type
		// Check if the user is an admin
		if userType != "admin" && userTypeCtx != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this route"})
			c.Abort()
			return
		}

		userId, exist := c.Get("user_id")
		if !exist || (userId != claims.UserID) {
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
