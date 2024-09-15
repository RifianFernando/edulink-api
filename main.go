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
var ALLOW_ORIGIN = os.Getenv("ALLOW_ORIGIN")

func init() {
	connections.LoadEnvVariables()
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")

	var isDevMode bool
	if strings.Contains(ALLOW_ORIGIN, "localhost") {
		isDevMode = true
	} else {
		isDevMode = false
	}
	// store.Options.Domain = os.Getenv("ALLOW_ORIGIN")
	store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   60, // 8 hours
		SameSite: http.SameSiteNoneMode,
		Secure:   isDevMode,
		Domain:   ALLOW_ORIGIN,
	}
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

	err := r.SetTrustedProxies(
		lib.HandleTrustedProxies(ALLOW_ORIGIN),
	)
	lib.HandleError(err, "Failed to set trusted proxies")

	routes.Route(r)

	err = r.Run()
	lib.HandleError(err, "Failed to serve the server")

	return r
}
