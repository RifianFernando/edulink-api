package middleware

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/gin-gonic/gin"
)

func AlreadyLoggedIn() gin.HandlerFunc {
	// validate access token
	return func(c *gin.Context) {
		accessToken, err := helper.GetCookieValue(c, "access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			c.Abort()
			return
		}

		// Set claims in the context
		c.Set("user_id", claims.UserID)
		c.Set("user_type", claims.User_type)
		c.Set("user_name", claims.UserName)

		c.Next()
	}
}

func IsNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		accessToken := helper.GetAuthTokenFromHeader(authHeader)
		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil {
			c.Next()
			return
		}

		// Set claims in the context
		c.Set("user_id", claims.UserID)
		c.Set("user_type", claims.User_type)
		c.Set("user_name", claims.UserName)

		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are already logged in"})
		c.Abort()
	}
}
