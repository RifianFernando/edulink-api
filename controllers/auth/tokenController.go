package controllers

import (
	"net/http"

	"github.com/edulink-api/config"
	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	// Retrieve the refresh token from the cookie
	refreshToken, err := helper.GetCookieValue(c, "token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
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

	c.JSON(http.StatusOK, gin.H{
		"access_token": newToken,
		"expires_in":   3600 * 24,
		"path":         "/",
		"domain":       config.ParsedDomain,
		"secure":       config.IsProdMode,
		"httponly":     true,
	})
}

func ValidateAccessToken(c *gin.Context) {
	// Retrieve the refresh token from the cookie
	accessToken, err := helper.GetCookieValue(c, "access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
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
