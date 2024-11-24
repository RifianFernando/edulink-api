package controllers

import (
	"net/http"

	"github.com/edulink-api/config"
	"github.com/edulink-api/helper"
	"github.com/edulink-api/request/auth"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		var req request.InsertLoginRequest
		req.UserEmail = email
		req.UserPassword = password

		// Bind and validate the request body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Authenticate the user
		user, userType := helper.Authenticate(req.UserEmail, req.UserPassword)
		if userType == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
			return
		}

		// Generate access token and refresh token
		accessToken, refreshToken, err := helper.GenerateToken(user, userType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		// Save session details in the database
		IpAddress := c.ClientIP()
		UserAgent := c.Request.UserAgent()
		errMsg := helper.InsertSession(refreshToken, user.UserID, IpAddress, UserAgent)

		// TODO: the best practice is if the token already exists, should request the verification method for user for missing the refresh token or access token
		if errMsg == "the refresh token already exists" {
			accessToken, refreshToken, err = helper.UpdateSession(refreshToken, user.UserID, IpAddress, UserAgent)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error()},
				)
				return
			}
		}

		// Set the refresh token in an HttpOnly cookie (valid for 1 day)
		// c.SetCookie("token", refreshToken, 3600*24*7, "/", config.ParsedDomain, config.IsProdMode, true)     // HttpOnly = true
		// c.SetCookie("access_token", accessToken, 3600*24, "/", config.ParsedDomain, config.IsProdMode, true) // HttpOnly = false
		cookie1 := http.Cookie{
			Name:  "token",
			Value: refreshToken,
			MaxAge: 3600 * 24 * 7,
			Path:  "/",
			Domain: config.ParsedDomain,
			Secure: config.IsProdMode,
			HttpOnly: true,
			SameSite: config.SameSite,
		}
		http.SetCookie(c.Writer, &cookie1)

		cookie2 := http.Cookie{
			Name:  "access_token",
			Value: accessToken,
			MaxAge: 3600 * 24 * 7,
			Path:  "/",
			Domain: config.ParsedDomain,
			Secure: config.IsProdMode,
			HttpOnly: true,
			SameSite: config.SameSite,
		}
		http.SetCookie(c.Writer, &cookie2)
		// Return success message and send the access token in the response body (optional)
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	}
}
