package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	error := router.Run(":9808")
	if error != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", error))
	}
}