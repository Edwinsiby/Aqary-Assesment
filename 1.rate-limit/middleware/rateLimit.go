package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IPRequestCount struct {
	Requests int
	LastTime int64
}

var ipRates = map[string]*IPRequestCount{}

func RateLimitCheck(c *gin.Context) {
	ip := c.ClientIP()

	now := time.Now().Unix()
	tracker, ok := ipRates[ip]

	if !ok || now-tracker.LastTime >= 10 {
		ipRates[ip] = &IPRequestCount{Requests: 1, LastTime: now}
	} else if tracker.Requests < 5 {
		tracker.Requests++
		ipRates[ip] = tracker
	} else {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "Too many requests",
		})
		c.Abort()
		return
	}
}
