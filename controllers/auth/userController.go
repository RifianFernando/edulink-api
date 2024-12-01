package controllers

import (
	"net/http"

	"github.com/edulink-api/helper"
	"github.com/edulink-api/models"
	"github.com/gin-gonic/gin"
)

func GetUserType(c *gin.Context) {
	userType, exist := c.Get("user_type")
	if !exist {
		uid, isExist := c.Get("user_id")
		if !isExist {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User type not found",
			})
			return
		}

		var user models.User
		user.UserID = uid.(int64)
		userType = helper.GetUserTypeByUID(user)
		if userType == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User type not found",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"user_type": userType,
	})
}
