package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/lib"
)

// Cors is a function to enable cors
func Cors() gin.HandlerFunc {
	// return the cors config
	return cors.New(cors.Config{
		AllowOrigins:     []string{lib.GetEnvValue("ALLOW_ORIGIN")},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			return origin == lib.GetEnvValue("ALLOW_ORIGIN")
		},
	})
}

