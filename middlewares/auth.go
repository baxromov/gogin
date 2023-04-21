package middlewares

import (
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function that checks if the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated
		if _, err := c.Cookie("access_token"); err != nil {
			// If the access token cookie is missing or invalid, redirect the user to the login page
			c.Redirect(404, "/login")
			return
		}

		// If the user is authenticated, call the next middleware/handler in the chain
		c.Next()
	}
}
