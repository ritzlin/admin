package main

import (
	"fmt"
	"log"
	"server/api"
	"server/global"

	"github.com/gin-gonic/gin"
)

func main() {
	global.Init()

	engine := gin.Default()
	api.InitRouter(engine)
	err := engine.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
	if err != nil {
		log.Fatal(err)
		return
	}
}
