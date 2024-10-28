package controllers

import (
	"net/http"
	"os"

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
			c.JSON(http.StatusNotFound, gin.H{"error": "user not registered"})
			return
		}

		// Generate a password reset token
		resetTokenLink, err := helper.GenerateResetPasswordToken(user.UserID, user.UserEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating reset token"})
			return
		}

		// TODO: GET the real domain without changing the code
		resetTokenLink = os.Getenv("ALLOW_ORIGIN") + "/auth/reset-password?token=" + resetTokenLink + "&email=" + user.UserEmail

		// Send the reset token to the user's email
		// err = helper.SendResetTokenEmail(user.UserEmail, resetToken)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{
			"message": "Reset token sent to your email",
			"reset_token_link":   resetTokenLink,
		})
	}
}

func ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.ResetPasswordRequest
		req.Token = c.PostForm("token")
		req.Email = c.PostForm("email")
		req.NewPassword = c.PostForm("password")
		req.ConfirmPassword = c.PostForm("password_confirmation")

		// Bind and validate the request body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the password
		if err := req.ValidatePasswords(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match"})
			return
		}

		// verify the reset token
		claims, msg := helper.ValidateResetPasswordToken(req.Token)
		if msg != "" || claims.UserEmail != req.Email {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		// Check if the email exists in the database
		user, err := helper.GetUserByEmail(req.Email)
		if err != nil || user.UserID == 0 || user.UserEmail != req.Email || user.UserID != claims.UserID {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not registered"})
			return
		}

		// Update the user's password
		// err = helper.UpdateUserPassword(user.UserID, req.Password)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password"})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{
			"token": req.Token,
			"email": req.Email,
		})
	}
}
