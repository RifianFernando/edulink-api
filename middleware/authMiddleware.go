package middleware

import (
	"github.com/gin-gonic/gin"
	// "github.com/skripsi-be/models"
)

func myAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(401, gin.H{"error middleware": "Unauthorized"})
            c.Abort()
            return
        }
        // Validate the token (implementation depends on your auth mechanism)
        // ...

        c.Next()
    }
}
