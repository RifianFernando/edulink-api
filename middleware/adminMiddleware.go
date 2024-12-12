package middleware

import (
	"net/http"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/database/user"
	"github.com/edulink-api/helper"
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := helper.GetCookieValue(c, "access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			c.Abort()
			return
		}

		// Check if the user is an admin
		isAdmin := user.CheckUserRole(claims.User_type, user.Admin)
				
		if !isAdmin && user.ValidateUserRoleCtx(c, user.Admin) {
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
