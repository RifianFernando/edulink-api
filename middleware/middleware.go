package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func getUserTypeFromContext(c *gin.Context) (string, error) {
	userTypeCtx, exist := c.Get("user_type")
	if !exist {
		return "", errors.New("user type not found")
	}
	return userTypeCtx.(string), nil
}
