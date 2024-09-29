package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/helper"
	"github.com/skripsi-be/models"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, valid := helper.GetClaimsToken(c)
		if !valid {
			c.Abort()
			return
		}

		var admin models.Admin
		fmt.Println("claims.getadminbyid: ", admin.GetAdminByUserID(claims.UserID))
		if claims.User_type == "admin" && admin.GetAdminByUserID(claims.UserID) == nil {
			c.Set("user_id", claims.UserID)
			c.Set("user_type", claims.User_type)
			c.Set("user_name", claims.UserName)
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this route"})
		c.Abort()
	}
}
