package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/skripsi-be/config"
	"github.com/skripsi-be/helper"
	"github.com/skripsi-be/models"
)

// AuthHandler handles authentication and user session validation.
func AuthHandler(isLoggedIn bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// Attempt to validate the refresh token from the cookie
			refreshToken, err := c.Cookie("token")
			if err != nil {
				if isLoggedIn {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
					c.Abort()
				} else {
					c.Next()
				}
				return
			}

			claims, msg := helper.ValidateRefreshToken(refreshToken)
			if msg != "" || claims == nil {
				if isLoggedIn {
					c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
					c.Abort()
				} else {
					c.Next()
				}
				return
			}

			// Update session and generate new tokens
			user := models.User{UserID: claims.UserID}
			ipAddress := c.ClientIP()
			userAgent := c.Request.UserAgent()
			newToken, newRefreshToken, err := helper.UpdateSession(refreshToken, user.UserID, ipAddress, userAgent)

			if err != nil {
				if isLoggedIn {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update session", "msg": err.Error()})
				}
				c.Next()
				return
			}
			// Set the refresh token in an HttpOnly cookie (valid for 1 day)
			c.SetCookie("token", newRefreshToken, 3600*24*7, "/", config.ParsedDomain, config.IsProdMode, true)

			if isLoggedIn {
				c.JSON(http.StatusOK, gin.H{
					"access_token":  newToken,
					"refresh_token": newRefreshToken,
				})
				c.Abort()
				return
			} else {
				c.Next()
			}
		}

		// Split the header to get the token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			if isLoggedIn {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
				c.Abort()
			} else {
				c.Next()
			}
			return
		}

		accessToken := parts[1]
		if accessToken == "" {
			if isLoggedIn {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
				c.Abort()
			} else {
				c.Next()
			}
			return
		}

		// Validate the access token
		claims, msg := helper.ValidateToken(accessToken, "access_token")
		if msg != "" {
			if isLoggedIn {
				c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
				c.Abort()
			} else {
				c.Next()
			}
			return
		}

		// Set claims in the context
		c.Set("user_id", claims.UserID)
		c.Set("user_type", claims.User_type)
		c.Set("user_name", claims.UserName)

		// If not logged in, check the user ID and respond accordingly
		if !isLoggedIn {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are already logged in"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AlreadyLoggedIn checks if a user is already logged in.
func AlreadyLoggedIn() gin.HandlerFunc {
	return AuthHandler(true)
}

// IsNotLoggedIn checks if a user is not logged in.
func IsNotLoggedIn() gin.HandlerFunc {
	return AuthHandler(false)
}
