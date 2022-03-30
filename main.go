package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string = ":" + os.Getenv("PORT")
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(PORT)
}