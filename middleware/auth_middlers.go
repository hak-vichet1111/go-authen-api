package middleware

import (
	"go-authentication-api/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// 1. Try getting token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			// The token should be prefixed with "Bearer"
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
				tokenString = tokenParts[1]
			}
		}

		// 2. If not in header, try getting from query parameter (for browser testing)
		if tokenString == "" {
			tokenString = c.Query("token")
		}

		// Debug log to see what's happening
		if tokenString != "" {
			log.Println("Token found:", tokenString[:10]+"...")
		} else {
			log.Println("No token found in Header or Query param")
		}

		// 3. If still empty, return unauthorized
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header or token query parameter is required"})
			c.Abort()
			return
		}

		// Verify the token
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			log.Println("Invalid token:", err)
			c.Abort()
			return
		}

		// Set the user ID in the context
		c.Set("user_id", claims["user_id"])

		c.Next()
	}
}
