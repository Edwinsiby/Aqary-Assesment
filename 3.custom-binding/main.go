package main

import (
	"custom/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("bind", handlers.CustomBinding)

	router.Run()
}
