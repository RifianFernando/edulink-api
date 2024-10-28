package main

import (
	"net/http"

	"github.com/edulink-api/config"
	"github.com/edulink-api/connections"
	"github.com/edulink-api/lib"
	"github.com/edulink-api/routes"
	"github.com/gin-gonic/gin"
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
	r := gin.New() // Creates a bare instance without Logger or Recovery
	r.Use(config.SetSecurityHeaders())
	r.Use(config.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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
