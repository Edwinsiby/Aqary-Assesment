package main

import (
	"limit/handlers"
	"limit/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("limit/check", middleware.RateLimitCheck, handlers.Test)

	router.Run(":8080")
}
