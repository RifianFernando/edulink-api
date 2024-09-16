package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	// "github.com/gin-gonic/gin"
	// "github.com/skripsi-be/models"
)

func IsLoggedIn(c *gin.Context) {
	var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	fmt.Println("auth middleware running")
	session, _ := store.Get(c.Request, "session")
	fmt.Println("session:", session)
	_, ok := session.Values["user"]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	fmt.Println("middleware done")
	c.JSON(http.StatusOK, gin.H{"message": "Authorized"})
	c.Next()
}
