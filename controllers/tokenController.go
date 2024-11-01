package controllers

import (
	"fmt"
	"net/http"

	"github.com/edulink-api/config"
	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	"github.com/gin-gonic/gin"
)

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		refreshToken, err := c.Cookie("token");
		if refreshToken == "" || err != nil {
			fmt.Println("No refresh token provided")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No refresh token provided",
			})
			return
		}

		claims, msg := helper.ValidateRefreshToken(refreshToken)
		if msg != "" || claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": msg,
			})
			return
		}

		// Update session and generate new tokens
		user := models.User{UserID: claims.UserID}
		ipAddress := c.ClientIP()
		userAgent := c.Request.UserAgent()
		newToken, newRefreshToken, err := helper.UpdateSession(refreshToken, user.UserID, ipAddress, userAgent)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update session", "msg": err.Error()})
			return
		}
		// Set the refresh token in an HttpOnly cookie (valid for 1 day)
		c.SetCookie("token", newRefreshToken, 3600*24*7, "/", config.ParsedDomain, config.IsProdMode, true)
		c.SetCookie("access_token", newToken, 3600*24, "/", config.ParsedDomain, config.IsProdMode, true)

		// for name, values := range c.Writer.Header() {
		// 	fmt.Printf("Response Header %s: %v\n", name, values)
		// }

		c.JSON(http.StatusOK, gin.H{
			"access_token": newToken,
			"expires_in":   3600 * 24,
			"path":         "/",
			"domain":       config.ParsedDomain,
			"secure":       config.IsProdMode,
			"httponly":     true,
		})
	}
}

func ValidateAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token");

		if accessToken == "" || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No access token provided",
			})
			return
		}

		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" || claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": msg,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Access token is valid",
		})
	}
}
