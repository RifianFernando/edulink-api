package controllers

import (
	"net/http"

	"github.com/edulink-api/config"
	"github.com/edulink-api/helper"
	request "github.com/edulink-api/request/auth"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
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
	if len(userType) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Generate access token and refresh token
	accessToken, refreshToken, err := helper.GenerateToken(user, helper.GetUserTypeByPrivilege(user))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// Save session details in the database
	IpAddress := c.ClientIP()
	UserAgent := c.Request.UserAgent()
	errMsg := helper.InsertSession(refreshToken, user.UserID, IpAddress, UserAgent)

	// the best practice is if the token already exists, should request the verification method for user for missing the refresh token or access token
	if errMsg == "the refresh token already exists" {
		accessToken, refreshToken, err = helper.UpdateSession(refreshToken, user.UserID, IpAddress, UserAgent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()},
			)
			return
		}
	}

	// set the output for the user session login info
	type userDTO struct {
		UserID int64    `json:"UserID"`
		Name   string   `json:"name"`
		Email  string   `json:"email"`
		Role   []string `json:"role"`
	}
	var UserDTO userDTO
	UserDTO.UserID = user.UserID
	UserDTO.Name = user.UserName
	UserDTO.Email = user.UserEmail
	UserDTO.Role = userType

	// Set the refresh token in an HttpOnly cookie (valid for 1 day)
	c.SetCookie("access_token", accessToken, 3600, "/", config.ParsedDomain, config.IsProdMode, true)
	c.SetCookie("token", refreshToken, 3600*24*7, "/", config.ParsedDomain, config.IsProdMode, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    UserDTO,
	})
}
