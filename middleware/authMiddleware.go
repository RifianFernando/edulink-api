package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/helper"
)

func IsLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, valid := helper.GetClaimsToken(c)
		if !valid {
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_type", claims.User_type)
		c.Set("user_name", claims.UserName)

		c.Next()
	}
}

func IsNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken, err := c.Cookie("token")

		if err != nil {
			c.Next()
			return
		}

		if clientToken == "" {
			c.Next()
			return
		}

		claims, msg := helper.ValidateToken(clientToken)
		userID, exist := c.Get("user_id")
		if msg != "" && claims.UserID == 0 && !exist && userID == nil {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are already logged in"})
		c.Abort()
	}
}
