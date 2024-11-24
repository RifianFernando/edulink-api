package middleware

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func AdminStaffOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessTokenHttp, _ := c.Request.Cookie("access_token")
		//  gett user from access token
		accessToken, err := c.Cookie("access_token")
		if (accessToken == "" || err != nil) && accessTokenHttp == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not found"})
			c.Abort()
			return
		} else if accessToken == "" {
			accessToken = accessTokenHttp.Value
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
		if (userType != "admin" && userTypeCtx != "admin") && (userType != "staff" && userTypeCtx != "staff") {
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

		var isAdmin, isStaff bool
		// Check if the admin exists
		var admin = models.Admin{
			UserID: userId.(int64),
		}
		err = admin.GetAdminByUserID()
		if err == nil && admin.UserID != 0 && admin != (models.Admin{}) {
			isAdmin = true
		}

		// check if the staff exists
		var staff = models.Staff{
			UserID: userId.(int64),
		}
		err = staff.GetStaffByModel()
		if err == nil && staff.UserID != 0 && staff != (models.Staff{}) {
			isStaff = true
		}

		if !isAdmin && !isStaff {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}
