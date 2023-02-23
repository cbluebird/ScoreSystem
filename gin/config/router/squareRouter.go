package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/squareController"
)

func squareRouterInit(r *gin.RouterGroup) {
	square := r.Group("/square")
	{
		square.POST("/pulish", squareController.PulishTopic)
		square.GET("/delete", squareController.DeleteTopic)
		square.GET("/iquire", squareController.InquireTopic)
	}
}
