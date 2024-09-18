package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/config"
)

func IsLoggedIn(c *gin.Context) {
	fmt.Println("auth middleware running")

	// Get session from the request
	session, err := config.Store.Get(c.Request, "session")
	if err != nil {
		fmt.Printf("Failed to get session: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		c.Abort()
		return
	}

	// Check if the session contains a user
	userId, ok := session.Values["userId"] // Ensure "userId" matches what is set during login
	if !ok || userId == nil {
		fmt.Println("Unauthorized access - userId not found in session")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// If user is authenticated, proceed
	c.Next()
}
