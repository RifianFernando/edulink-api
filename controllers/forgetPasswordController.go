package controllers

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/request"
	"github.com/gin-gonic/gin"
)

func ForgetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")

		var req request.InsertForgetPasswordRequest
		req.UserEmail = email

		// Bind and validate the request body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the email exists in the database
		user, err := helper.GetUserByEmail(req.UserEmail)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Generate a password reset token
		// resetToken, err := helper.GenerateResetToken(user.UserID)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating reset token"})
		// 	return
		// }

		// Send the reset token to the user's email
		// err = helper.SendResetTokenEmail(user.UserEmail, resetToken)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{
			"message": "Reset token sent to your email",
			"user":    user,
		})
	}
}
