package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/studentController"
	"scoresystem/app/services/studentService"

	"scoresystem/app/midwares"
)

func studentRouterInit(r *gin.RouterGroup) {
	student := r.Group("/student")
	{
		student.POST("/add/application", studentController.Add, midwares.CheckLogin)
		student.GET("/get/grade", studentController.GetAll, midwares.CheckLogin)
		student.GET("/get/detailscore", studentController.GetDetail, midwares.CheckLogin)
		student.GET("/inquire", studentService.InquireApplication, midwares.CheckLogin)
		student.POST("/pulishSuggestion", studentController.PulishSuggest, midwares.CheckLogin)
		student.POST("/updateSuggestion", studentController.UpdateSuggestion, midwares.CheckLogin)
		student.GET("/iquireSuggestion", studentController.InquireSuggestion, midwares.CheckLogin)
		student.GET("/deleteSuggestion", studentController.DeleteSuggestion, midwares.CheckLogin)
		student.POST("/creatappeal", studentController.CreatAppeal, midwares.CheckLogin)
	}
}
