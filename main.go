package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/config"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/routes"
)

func init() {
	connections.LoadEnvVariables()
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
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

	err := r.Run()
	lib.HandleError(err, "Failed serve the server")
}
