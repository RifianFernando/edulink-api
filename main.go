package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/config"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/routes"
)

func init() {
	// Load environment variables
	connections.LoadEnvVariables()
	config.InitializeSessionStore()

	// Initialize the database connection
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

func main() {
	r := setupRouter()

	// Start the server and handle potential errors
	err := r.Run()
	lib.HandleError(err, "Failed to serve the server")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(config.SetSecurityHeaders())
	r.Use(config.Cors())
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		cookieResult, err := c.Cookie("token")
		if err != nil {
			cookieResult = "No cookie"
		}

		currentTime := lib.GetTimeNow().Format("2006-01-02 15:04:05 MST")
		c.JSON(http.StatusOK, gin.H{
			"current_time": currentTime,
			"cookie":       cookieResult,
		})
	})

	r.ForwardedByClientIP = true

	// Set trusted proxies or use a real proxy list for production
	err := r.SetTrustedProxies(nil)
	lib.HandleError(err, "Failed to set trusted proxies")

	routes.Route(r)

	return r
}
