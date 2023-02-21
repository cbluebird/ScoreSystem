package router

import (
	"scoresystem/app/controllers/squareController"
	"scoresystem/app/midwares"

	"github.com/gin-gonic/gin"
)

func squareRouterInit(r *gin.RouterGroup) {
	square := r.Group("/square")
	{
		square.POST("/pulish", midwares.CheckLogin, squareController.PulishTopic)
		square.GET("/delete", midwares.CheckLogin, squareController.DeleteTopic)
		square.GET("/iquire", midwares.CheckLogin, squareController.InquireTopic)
	}
}
