package middleware

import (
	"fmt"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/helper"
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
			fmt.Println("Error getting user type from context: ", err)
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

// need to commented this because if the roles is admin and teacher the user will be considered as admin but the user is also a teacher and should get the access of the code
func OnlyTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := helper.GetCookieValue(c, "access_token")
		if err != nil {
			res.AbortUnauthorized(c)
			return
		}

		userTypeCtx, err := getUserTypeFromContext(c)
		if err != nil {
			fmt.Println("Error getting user type from context: ", err)
			res.AbortUnauthorized(c)
			return
		}

		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil {
			res.AbortUnauthorized(c)
			return
		}

		if userTypeCtx != "teacher" {
			userType := helper.GetUserTypeByUID(
				models.User{
					UserID: claims.UserID,
				},
			)
			for _, role := range userType {
				if role == "teacher" {
					c.Next()
					return
				}
			}
			res.AbortUnauthorized(c)
			return
		}

		if isAdminOrStaff(claims.UserID) {
			res.AbortUnauthorized(c)
			return
		}

		c.Next()
	}
}
