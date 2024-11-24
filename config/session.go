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
		ParsedDomain = ""
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
		Path:     "/", // This should be the same as the router group base path
	}

	fmt.Println("Is in Production mode:", IsProdMode)
	fmt.Println("maxAge:", Store.Options.MaxAge)
	fmt.Println("Parsed Domain:", ParsedDomain)
	fmt.Println("Same Site:", SameSite)
}

func extractDomain(fullUrl string) string {
	fullUrl = strings.TrimPrefix(fullUrl, "https://")
	fullUrl = strings.TrimPrefix(fullUrl, "http://")
	return fullUrl
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
