package middleware

import (
	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

// IsTeacherHomeRoom middleware checks if the user is authorized as a home room teacher.
func IsTeacherHomeRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := helper.GetCookieValue(c, "access_token")
		if err != nil {
			res.AbortUnauthorized(c)
			return
		}

		userTypeCtx, err := getUserTypeFromContext(c)
		if err != nil {
			res.AbortUnauthorized(c)
			return
		}

		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil {
			res.AbortUnauthorized(c)
			return
		}

		if isAdminOrStaff(claims.UserID) {
			c.Next()
			return
		}

		if isTeacherHomeRoom(claims.UserID, claims.User_type, userTypeCtx) {
			c.Next()
			return
		}

		res.AbortUnauthorized(c)
	}
}

func isTeacherHomeRoom(userID int64, userType string, userTypeCtx interface{}) bool {
	if userType != "teacher" || userTypeCtx != "teacher" {
		return false
	}

	var teacher models.Teacher
	teacher.UserID = userID
	if err := teacher.GetTeacherByModel(); err != nil {
		return false
	}

	var className models.ClassName
	className.TeacherID = teacher.TeacherID
	classes, err := className.GetHomeRoomTeacherByTeacherID()
	if err != nil {
		return false
	}

	if len(classes) == 0 {
		return false
	}

	return true
}
