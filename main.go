package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/skripsi-be/config"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/routes"
)

var session = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func init() {
	session.Options.HttpOnly = true
	session.Options.SameSite = http.SameSiteLaxMode
	// if (lib.GetEnvValue("ENV") == "production") {
	// 	session.Options.Secure = true
	// }
	connections.LoadEnvVariables()
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

func main() {
	// use router
	router()
}

func router() *gin.Engine {
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

	return r
}
