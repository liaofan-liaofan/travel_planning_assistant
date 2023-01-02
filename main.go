package main

import "github.com/gin-gonic/gin"

//go:generate go mod tidy -go=1.19
//go:generate go mod download

func main() {
	r := gin.Default()
	r.GET("api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "hello",
		})
	})
	r.Run(":8888")
}
