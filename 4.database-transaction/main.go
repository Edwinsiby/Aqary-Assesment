package main

import (
	"db/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("user/data", handlers.AddStudentData)
}
