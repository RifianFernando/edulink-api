package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/skripsi-be/config"
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/lib"
	"github.com/skripsi-be/routes"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func init() {
	// store.Options.HttpOnly = true
	// store.Options.MaxAge = 60
	// store.Options.SameSite = http.SameSiteNoneMode
	var isDevMode bool
	if strings.Contains(os.Getenv("ALLOW_ORIGIN"), "localhost") {
		// store.Options.Secure = false // For local development; set to true in production
		isDevMode = true
	} else {
		// store.Options.Secure = true
		isDevMode = false
	}
	// store.Options.Domain = os.Getenv("ALLOW_ORIGIN")
	store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   60, // 8 hours
		SameSite: http.SameSiteNoneMode,
		Secure:   isDevMode,
		Domain:   os.Getenv("ALLOW_ORIGIN"),
	}

	connections.LoadEnvVariables()
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")
}

func main() {
	router()
}

func router() *gin.Engine {
	r := gin.Default()
	r.Use(config.Cors())
	r.GET("/", func(c *gin.Context) {
		currentTime := time.Now().Format("2006-01-02 15:04:05 MST")
		c.JSON(200, gin.H{
			"current_time": currentTime,
		})
	})
	r.ForwardedByClientIP = true
	r.SetTrustedProxies(
		[]string{
			os.Getenv("ALLOW_ORIGIN"),
		},
	)

	routes.Route(r)

	err := r.Run()
	lib.HandleError(err, "Failed to serve the server")

	return r
}
