package middlewares

//goland:noinspection ALL
import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware is a middleware function that logs each request
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the start time of the request
		start := time.Now()

		// Call the next middleware/handler in the chain
		c.Next()

		// Get the end time of the request
		end := time.Now()

		// Get the duration of the request
		duration := end.Sub(start)

		// Log the request
		fmt.Printf("%s %s %s %s %d %s\n",
			end.Format("2006-01-02 15:04:05"),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration.String(),
		)
	}
}
