package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"scoresystem/app/midwares"
	"scoresystem/config/config"
	"scoresystem/config/router"
	"scoresystem/config/session"
	"scoresystem/config/startConfig"
)

func main() {
	startConfig.Init()
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	session.Init(r)
	router.Init(r)
	err := r.Run(":" + config.Config.GetString("server.port"))
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
