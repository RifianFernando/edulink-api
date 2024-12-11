package controllers

import (
	"net/http"

	"github.com/edulink-api/config"
	"github.com/edulink-api/helper"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	// Retrieve the refresh token from the cookie
	refreshToken, err := helper.GetCookieValue(c, "token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Retrieve the accees token from the cookie
	accessToken, err := helper.GetCookieValue(c, "access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Delete the refresh token from the server
	isDeleted, msg := helper.DeleteToken(accessToken, refreshToken)

	if !isDeleted {
		// Handle error if token deletion fails
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	// Remove the refresh token cookie
	c.SetCookie("token", "", -1, "/", config.ParsedDomain, config.IsProdMode, true) // Clear the cookie
	c.SetCookie("access_token", "", -1, "/", config.ParsedDomain, config.IsProdMode, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
