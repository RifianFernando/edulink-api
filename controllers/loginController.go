package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/helper"
	"github.com/skripsi-be/request"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		var req request.InsertLoginRequest
		req.UserEmail = email
		req.UserPassword = password

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, userType := helper.Authenticate(req.UserEmail, req.UserPassword)
		if userType == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials when authenticating",
			})
			return
		}

		token, refreshToken, err := helper.GenerateToken(user, userType)
		if err != nil {
			fmt.Sprintln("Error generating token: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var IpAddress = c.ClientIP()
		var UserAgent = c.Request.UserAgent()
		err = helper.UpdateSessionTable(token, refreshToken, user.UserID, IpAddress, UserAgent)
		if err != nil {
			fmt.Sprintln("Error updating session table: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// c.SetCookie("token", token, 3600*24, "/", config.ParsedDomain, config.IsProdMode, true) // true for HttpOnly
		c.SetCookie("token", token, 3600, "/", "", false, true) // Adjust as needed

		// c.JSON(http.StatusOK, gin.H{"message": "Login success"})
		c.JSON(http.StatusOK, gin.H{
			"message": "Login success",
			"token":   token,
		})
	}
}
