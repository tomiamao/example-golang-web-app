package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong pong pong ------ NEW MASTER",
		})
	})
	r.Run(fmt.Sprintf(":%s", port))
}
