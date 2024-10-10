package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/config"
	"github.com/skripsi-be/helper"
)

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the refresh token from the cookie
		refreshToken, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token not found",
			})
			return
		}

		// Delete the refresh token from the server (if applicable)
		isDeleted, msg := helper.DeleteToken(refreshToken)

		if !isDeleted {
			// Handle error if token deletion fails
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": msg,
			})
			return
		}

		// Remove the refresh token cookie
		c.SetCookie("token", "", -1, "/", config.ParsedDomain, config.IsProdMode, true) // Clear the cookie

		c.JSON(http.StatusOK, gin.H{
			"message": "Logout successful",
		})
	}
}
