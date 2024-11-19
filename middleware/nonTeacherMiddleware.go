package middleware

import (
	"fmt"
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	"github.com/gin-gonic/gin"
)

var (
	forbidden = "Forbidden"
)

func IsTeacherHomeRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		//  gett user from access token
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": forbidden})
			c.Abort()
			return
		}

		// Get the user type from the context
		userTypeCtx, exist := c.Get("user_type")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": forbidden})
			c.Abort()
			return
		}

		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": forbidden})
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
			// teacher.UserID = claims.UserID
			err := teacher.GetTeacherByModel()
			if(err != nil){
				c.Abort()
				return
			}
			fmt.Println("Teacher: ", teacher)
			c.Abort()
			return
		}
		c.Abort()
	}
}
