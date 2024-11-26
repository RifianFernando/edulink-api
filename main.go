// package handler
package main

import (
	"net/http"

	"github.com/edulink-api/config"
	"github.com/edulink-api/connections"
	"github.com/edulink-api/lib"
	_ "github.com/edulink-api/request"
	"github.com/edulink-api/routes"
	"github.com/gin-gonic/gin"
)

// init function for environment setup
func init() {
	connections.LoadEnvVariables()
	config.InitializeSessionStore()

	// Initialize database connection
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

// Vercel requires an HTTP handler function that serves the API
// func Handler(w http.ResponseWriter, r *http.Request) {
//     // Set up the router (you don't need r.Run() in serverless)
//     app := setupRouter()

//	    // Use the router to handle the HTTP request
//	    app.ServeHTTP(w, r)
//	}
func main() {
	// Set up the router
	r := setupRouter()

	// Run the server
	r.Run() // listen and serve on
}

// Set up the Gin router
func setupRouter() *gin.Engine {
	r := gin.New() // Create a bare instance of Gin
	r.Use(config.Cors())
	r.Use(config.SetSecurityHeaders())
	r.Use(config.SessionMiddleware())
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

	// Forward the client's IP address
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies(nil) // Adjust for production as needed
	lib.HandleError(err, "Failed to set trusted proxies")

	// Register routes
	routes.Route(r)

	return r
}
