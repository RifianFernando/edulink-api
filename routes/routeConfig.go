package routes

import (
	"net/http"

	"github.com/edulink-api/config"
	"github.com/edulink-api/helper"
	"github.com/edulink-api/lib"
	"github.com/gin-gonic/gin"
)

// Set up the Gin router
func SetupRouter() *gin.Engine {
	r := gin.New() // Create a bare instance of Gin
	r.Use(config.Cors())
	r.Use(config.SetSecurityHeaders())
	r.Use(config.SessionMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		cookieResult, err := helper.GetCookieValue(c, "token")
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
	Route(r)

	return r
}
