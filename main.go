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
		url := URL{}
		c.BindJSON(&url)
		_, err := net.Dial("tcp", url.Name + ":80")
		if err != nil {
			c.JSON(200, gin.H{
				"message": "down",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "up",
			})
		}
	})
	r.Run(PORT)
}