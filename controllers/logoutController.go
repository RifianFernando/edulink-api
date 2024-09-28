package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/config"
	"github.com/skripsi-be/helper"
)

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken, err := c.Cookie("token")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Token not found",
			})
			return
		}

		isDeleted, msg := helper.DeleteToken(clientToken)

		if !isDeleted && msg != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": msg,
			})
			return
		}

		c.SetCookie("token", "", -1, "/", config.ParsedDomain, config.IsProdMode, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "Logout success",
		})
	}
}
