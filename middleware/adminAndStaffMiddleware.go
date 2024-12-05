package middleware

import (
	"errors"
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/database/models"
	"github.com/edulink-api/res"
	"github.com/gin-gonic/gin"
)

func AdminStaffOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := helper.GetCookieValue(c, "access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		userTypeCtx, err := getUserTypeFromContext(c)
		if err != nil || userTypeCtx == "" {
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

		if !isAuthorizedUser(claims.User_type, userTypeCtx) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this route"})
			c.Abort()
			return
		}

		userId, err := getUserIdFromContext(c, claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !isAdminOrStaff(userId) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": res.Forbidden})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}

func isAuthorizedUser(userType, userTypeCtx string) bool {
	return (userType == "admin" || userTypeCtx == "admin") || (userType == "staff" || userTypeCtx == "staff")
}

func getUserIdFromContext(c *gin.Context, expectedUserId int64) (int64, error) {
	userId, exist := c.Get("user_id")
	if !exist || (userId != expectedUserId) {
		return 0, errors.New("user id not found")
	}
	return userId.(int64), nil
}

func isAdminOrStaff(userId int64) bool {
	userType := helper.GetUserTypeByUID(
		models.User{
			UserID: userId,
		},
	)

	if userType == "admin" || userType == "staff" {
		return true
	}

	return false
}
