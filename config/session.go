package config

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
)

var (
	// Exported session store variable
	Store *sessions.CookieStore
	// Production mode flag
	IsProdMode bool
)

func InitializeSessionStore() {
	allowOrigin := os.Getenv("ALLOW_ORIGIN")
	sessionKey := os.Getenv("SESSION_KEY")
	var parsedDomain string

	if sessionKey == "" {
		panic("SESSION_KEY is not set in the environment")
	}

	if strings.Contains(allowOrigin, "localhost") {
		IsProdMode = false
		parsedDomain = ""
	} else {
		IsProdMode = true
		parsedDomain = extractDomain(allowOrigin)
	}

	Store = sessions.NewCookieStore([]byte(sessionKey))
	Store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   8 * 60 * 60, // 8 hours
		SameSite: http.SameSiteLaxMode,
		Secure:   IsProdMode,
		Domain:   parsedDomain,
		Path:     "/", // This should be the same as the router group base path
	}

	fmt.Println("Is in Production mode:", IsProdMode)
	fmt.Println("Parsed Domain:", parsedDomain)
}

// Helper function to extract base domain from a full URL
func extractDomain(fullUrl string) string {
	fullUrl = strings.TrimPrefix(fullUrl, "https://")
	fullUrl = strings.TrimPrefix(fullUrl, "http://")
	return fullUrl
}
