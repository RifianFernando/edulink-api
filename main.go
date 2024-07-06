package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/routes"
)

func init() {
	connections.LoadEnvVariables()
	connections.ConnecToDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{lib.GetEnvValue("ALLOW_ORIGIN")},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			return origin == lib.GetEnvValue("ALLOW_ORIGIN")
		},
	}))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.POST("/student/create", controllers.CreateStudent)

	// use router
	routes.Route(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
