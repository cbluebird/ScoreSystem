package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/studentController"
	"scoresystem/app/midwares"
)

func studentRouterInit(r *gin.RouterGroup) {
	student := r.Group("/student")
	{
		student.POST("/add/application", studentController.Add, midwares.CheckLogin)
	}
}
