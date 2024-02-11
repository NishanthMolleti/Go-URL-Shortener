package main

import (
	"fmt"

	"github.com/NishanthMolleti/go-url-shortner/handler"
	"github.com/NishanthMolleti/go-url-shortner/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})


	// Note that store initialization happens here
	storage.InitializeStore()

	error := router.Run(":9808")
	if error != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", error))

	}
}