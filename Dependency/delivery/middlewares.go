package delivery

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(c *gin.Context) {
	limit := tollbooth.NewLimiter(2, &limiter.ExpirableOptions{DefaultExpirationTTL: 2 * time.Second})
	httpError := tollbooth.LimitByRequest(limit, c.Writer, c.Request)
	if httpError != nil {
		c.JSON(httpError.StatusCode, gin.H{"error": httpError.Message})
		c.Abort()
	} else {
		c.Next()
	}
}
