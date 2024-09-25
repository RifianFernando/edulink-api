package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/config"
)

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", config.ParsedDomain, config.IsProdMode, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "Logout success",
		})
	}
}
