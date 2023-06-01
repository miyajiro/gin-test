package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello"})
	})

	router.Run("localhost:8000")
}