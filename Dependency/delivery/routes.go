package delivery

import (
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Handlers HandlerMethods
}

func InitRoutes(r *gin.Engine, h HandlerMethods, c *gin.Context) *gin.Engine {

	r.POST("limitcheck", RateLimiterMiddleware, h.RateLimitCheck)

	return r
}
