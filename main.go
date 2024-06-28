package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/routes"
)

func init() {
	connections.LoadEnvVariables()
	connections.ConnecToDB()
}

func main() {
	r := gin.Default()
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
