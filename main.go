package main

import (
	"fmt"
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

var (
	store        *sessions.CookieStore
	allowOrigin  string
	isProdMode   bool
	parsedDomain string
)

func init() {
	// Load environment variables
	connections.LoadEnvVariables()

	// Initialize the database connection
	err := connections.ConnecToDB()
	lib.HandleError(err, "Failed to connect db")

	// Read environment variables
	allowOrigin = os.Getenv("ALLOW_ORIGIN")
	sessionKey := os.Getenv("SESSION_KEY")

	// Determine if in production mode
	if strings.Contains(allowOrigin, "localhost") {
		isProdMode = false
		parsedDomain = ""
	} else {
		isProdMode = true
		parsedDomain = extractDomain(allowOrigin)
	}

	// Initialize the session store
	store = sessions.NewCookieStore([]byte(sessionKey))
	store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   8 * 60 * 60, // 8 hours
		SameSite: http.SameSiteNoneMode,
		Secure:   isProdMode, // Secure in production, not in dev
		Domain:   parsedDomain,
	}

	// Print debug information
	fmt.Println("Is in Production mode:", isProdMode)
	fmt.Println("Parsed Domain:", parsedDomain)
}

func main() {
	r := setupRouter()

	// Start the server and handle potential errors
	err := r.Run()
	lib.HandleError(err, "Failed to serve the server")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(config.Cors())

	r.GET("/", func(c *gin.Context) {
		currentTime := time.Now().Format("2006-01-02 15:04:05 MST")
		c.JSON(http.StatusOK, gin.H{
			"current_time": currentTime,
		})
	})

	r.ForwardedByClientIP = true

	// Set trusted proxies or use a real proxy list for production
	err := r.SetTrustedProxies(nil)
	lib.HandleError(err, "Failed to set trusted proxies")

	routes.Route(r)

	return r
}

// Example function to extract base domain from a full URL (e.g., "https://example.com")
func extractDomain(fullUrl string) string {
	fullUrl = strings.TrimPrefix(fullUrl, "https://")
	fullUrl = strings.TrimPrefix(fullUrl, "http://")
	// This is a placeholder, implement your domain extraction logic
	return fullUrl
}
