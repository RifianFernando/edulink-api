package config

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	Store        *sessions.CookieStore
	IsProdMode   bool
	ParsedDomain string
	SameSite     http.SameSite
)

func InitializeSessionStore() {
	allowOrigin := os.Getenv("ALLOW_ORIGIN")
	sessionKey := os.Getenv("APP_KEY")

	if sessionKey == "" {
		panic("APP_KEY is not set in the environment")
	}

	if strings.Contains(allowOrigin, "localhost") {
		IsProdMode = false
		ParsedDomain = "" // Or you can set it to "localhost" for testing
		gin.SetMode(gin.DebugMode)
		SameSite = http.SameSiteLaxMode
	} else {
		IsProdMode = true
		gin.SetMode(gin.ReleaseMode)
		ParsedDomain = extractDomain(allowOrigin)
		SameSite = http.SameSiteLaxMode
	}

	Store = sessions.NewCookieStore([]byte(sessionKey))
	Store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   7 * 24 * 60 * 60, // 7 days same as the token expiration
		SameSite: SameSite,
		Secure:   IsProdMode,
		Domain:   ParsedDomain,
		Path:     "/",
	}

	fmt.Println("Is in Production mode:", IsProdMode)
	fmt.Println("SameSite:", SameSite)
	fmt.Println("maxAge:", Store.Options.MaxAge)
	fmt.Println("Parsed Domain:", ParsedDomain)
}

func extractDomain(fullURL string) string {
	fullURL = strings.TrimPrefix(fullURL, "https://")
	fullURL = strings.TrimPrefix(fullURL, "http://")
	return fullURL
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := Store.Get(c.Request, "session") // Replace with a constant or config
		if err != nil {
			c.Abort()
			return
		}

		c.Set("session", session)

		defer func() {
			if err := sessions.Save(c.Request, c.Writer); err != nil {
				c.Abort()
				return
			}
		}()

		c.Next()
	}
}
