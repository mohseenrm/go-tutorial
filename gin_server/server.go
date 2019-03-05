package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// General routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Static files
	r.StaticFS("/files", http.Dir("static_files"))

	r.Run(":9000") // listen and serve on 0.0.0.0:8080
}
