package main

import (
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

type URL struct {
	Name string `json:"name"`
}

var (
	PORT string = ":" + os.Getenv("PORT")
)

func main() {
	r := gin.Default()

	r.POST("/up", func(c *gin.Context) {
		// check if url is up
		var url URL
		c.BindJSON(&url)
		_, err := net.Dial("tcp", url.Name)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "up",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "down",
			})
		}
	})
	r.Run(PORT)
}