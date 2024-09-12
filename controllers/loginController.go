package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/request"
)

func Login(c *gin.Context) {
	var request request.InsertLoginRequest

	// Bind the request JSON to the CreateStudentRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error should bind json": err.Error(),
		})
		return
	}

	// Validate the request
	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error when validate": err.Error(),
		})
		return
	}

	// Authenticate the user
	userID, err := Authenticate(request.UserEmail, request.UserPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"userID":  userID,
	})
}
