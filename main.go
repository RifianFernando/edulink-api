package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/routes"
	"github.com/skripsi-be/config"
)

func init() {
	connections.LoadEnvVariables()
	connections.ConnecToDB()
}

func main() {
	r := gin.Default()
	r.Use(config.Cors())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// use router
	routes.Route(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
