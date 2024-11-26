package middleware

import (
	"fmt"
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/gin-gonic/gin"
)

func AlreadyLoggedIn() gin.HandlerFunc {
	// validate access token
	return func(c *gin.Context) {
		// Get the Authorization header
		// authHeader := c.GetHeader("Authorization")
		// accessToken := helper.GetAuthTokenFromHeader(authHeader)
		accessTokenHttp, _ := c.Request.Cookie("access_token")
		accessToken, err := c.Cookie("access_token")
		if accessToken == "" || err != nil {
			if (accessTokenHttp.Value == "") {
				c.Abort()
				return
			}
			accessToken = accessTokenHttp.Value
		}

		fmt.Println("access token: ", accessToken)

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
