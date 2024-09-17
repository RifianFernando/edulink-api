package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func IsLoggedIn(c *gin.Context) {
	var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	fmt.Println("auth middleware running")

	// Get session from the request
	fmt.Println("Request cookies:", c)
	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		c.Abort()
		return
	}

	// Check if the session contains a user
	userId, ok := session.Values["userId"] // Ensure "userId" matches what is set during login
	if !ok || userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// If user is authenticated, proceed
	c.Next()
}
