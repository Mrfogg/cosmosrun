package main

import "github.com/gin-gonic/gin"

func main() {
	engin := gin.New()
	engin.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})
	engin.Run(":8888")
}
