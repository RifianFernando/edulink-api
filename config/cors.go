package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors is a function to enable cors
func Cors() gin.HandlerFunc {
	var allowOrigin = os.Getenv("ALLOW_ORIGIN")
	return cors.New(cors.Config{
		AllowOrigins:     []string{allowOrigin},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true, // This allows cookies to be sent
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			log.Printf("Origin: %s, AllowOrigin: %s", origin, allowOrigin)
			return strings.Contains(origin, allowOrigin)
		},
	})
}

func SetSecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")
		c.Next()
	}
}
