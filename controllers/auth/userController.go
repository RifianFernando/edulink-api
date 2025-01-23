package controllers

import (
	"net/http"

	"github.com/edulink-api/database/models"
	"github.com/edulink-api/database/user"
	"github.com/edulink-api/helper"
	"github.com/gin-gonic/gin"
)

func GetUserType(c *gin.Context) {
	userTypeArray := user.GetUserTypeFromCtx(c)
	if len(userTypeArray) == 0 {
		uid, isExist := c.Get("user_id")
		if !isExist {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User type not found",
			})
			return
		}

		var user models.User
		user.UserID = uid.(int64)
		userTypeArray = helper.GetUserTypeByUID(user)
		if len(userTypeArray) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User type not found",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"user_type": userTypeArray,
	})
}

func GetUserProfile(c *gin.Context) {
	uid, isExist := c.Get("user_id")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID not found",
		})
		return
	}

	userType := user.GetUserTypeFromCtx(c)
	if len(userType) == 0 {
		userType = helper.GetUserTypeByUID(models.User{UserID: uid.(int64)})
		if len(userType) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User type not found",
			})
			return
		}
	}

	var user models.User
	user.UserID = uid.(int64)
	result, err := user.GetUserProfile(userType)
	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profile": result,
	})
}
