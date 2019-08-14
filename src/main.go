package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
	"net/http"
)

func main(){
	config, err := toml.LoadFile("config/basic.toml")
	if err != nil{
		panic(err)
	}

	r := gin.Default()
	r.NoRoute(NoResponse)
	r.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"error": 0,
			"data": "Hi! Welcome to Apicon!",
		})
	})

	r.POST("/greeting", func(c *gin.Context) {
		name, exist := c.GetQuery("name")
		if !exist{
			name = "But I don't know who you are."
		}

		msg, _ := c.GetRawData()
		c.JSON(200, gin.H{
			"error": 0,
			"data": map[string]string{
				"say": "Hello! " + name,
				"message": string(msg),
			},
		})
	})
	_ = r.Run(config.Get("server.port").(string))
}

func NoResponse(c *gin.Context){
	c.JSON(http.StatusNotFound, gin.H{
		"error": 40400,
		"data":  "",
	})
}
