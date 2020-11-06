package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.NoRoute(NoResponse)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error": 0,
			"data":  "Hi! Welcome to Apicon!",
		})
	})

	r.POST("/greeting", func(c *gin.Context) {
		name, exist := c.GetQuery("name")
		if !exist {
			name = "But I don't know who you are."
		}

		msg, _ := c.GetRawData()
		c.JSON(200, gin.H{
			"error": 0,
			"data": map[string]string{
				"say":     "Hello! " + name,
				"message": string(msg),
			},
		})
	})
	_ = r.Run(":8080")
}

func NoResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": 40400,
		"data":  "",
	})
}
