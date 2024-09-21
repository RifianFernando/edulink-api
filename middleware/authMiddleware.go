package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/helper"
)

func IsLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// clientToken := c.Request.Header.Get("token")
		clientToken, err := c.Cookie("token")
		// clientToken, err := c.Request.Cookie("token")
		fmt.Println("All Cookies: ", c.Request.Cookies()) // Log all cookies for debugging

		fmt.Println("clientToken 1: ", clientToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found in cookies"})
			c.Abort()
			return
		}

		fmt.Println("clientToken: ", clientToken)

		claims, msg := helper.ValidateToken(clientToken)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_type", claims.User_type)
		c.Set("user_name", claims.UserName)

		c.Next()
	}
}
