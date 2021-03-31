package main

import (
	"dev.apicon.cn/sdk/service"
	"github.com/wuhan005/gadget"

	"github.com/gin-gonic/gin"
	log "unknwon.dev/clog/v2"
)

func main() {
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic(err)
	}

	svc := service.New("Greeting", 2)

	svc.Route().GET("/greeting", func(c *gin.Context) {
		c.JSON(gadget.MakeSuccessJSON(gin.H{
			"error": 0,
			"data":  "Hi! Welcome to Apicon!",
		}))
	})

	svc.Route().POST("/greeting", func(c *gin.Context) {
		var name string

		user, err := service.GetUser(c)
		if err == nil {
			name = user.Name
		} else {
			name = "But I don't know who you are."
		}

		msg, _ := c.GetRawData()
		c.JSON(gadget.MakeSuccessJSON(gin.H{
			"say":     "Hello " + name,
			"message": string(msg),
		}))
	})

	svc.Route().GET("/whoami", func(c *gin.Context) {
		user, err := service.GetUser(c)
		if err == nil {
			c.JSON(gadget.MakeSuccessJSON("Hello " + user.NickName))
			return
		}

		c.JSON(gadget.MakeSuccessJSON("Sorry, I don't know who you are."))
	})

	svc.Run()
}
