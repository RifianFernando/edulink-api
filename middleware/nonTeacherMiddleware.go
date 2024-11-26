package middleware

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func IsTeacherHomeRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		//  gett user from access token
		accessTokenHttp, _ := c.Request.Cookie("access_token")
		accessToken, err := c.Cookie("access_token")
		if (accessToken == "" || err != nil) && accessTokenHttp == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			c.Abort()
			return
		} else if accessToken == "" {
			accessToken = accessTokenHttp.Value
		}

		// Get the user type from the context
		userTypeCtx, exist := c.Get("user_type")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			c.Abort()
			return
		}

		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			c.Abort()
			return
		}

		userId, exist := c.Get("user_id")
		if !exist || (userId != claims.UserID) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User id not found"})
			c.Abort()
			return
		}

		userType := claims.User_type

		if (userType == "admin" && userTypeCtx == "admin") || (userType == "staff" && userTypeCtx == "staff") {
			c.Next()
			return
		}

		// Check if the user is a home room teacher
		if userType == "teacher" && userTypeCtx == "teacher" {
			var teacher models.Teacher
			teacher.UserID = claims.UserID
			err := teacher.GetTeacherByModel()
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
				c.Abort()
				return
			}

			var className models.ClassName
			className.TeacherID = teacher.TeacherID
			err = className.GetHomeRoomTeacherByTeacherID()
			if err != nil || className.ClassNameID == 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
				c.Abort()
				return
			}

			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
		c.Abort()
	}
}
