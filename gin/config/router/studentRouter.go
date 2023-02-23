package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/studentController"
	"scoresystem/app/services/studentService"
)

func studentRouterInit(r *gin.RouterGroup) {
	student := r.Group("/student")
	{
		student.POST("/add/application", studentController.Add)
		student.GET("/get/grade", studentController.GetAll)
		student.GET("/get/detailscore", studentController.GetDetail)
		student.GET("/inquire", studentService.InquireApplication)
		student.POST("/pulishSuggestion", studentController.PulishSuggest)
		student.POST("/updateSuggestion", studentController.UpdateSuggestion)
		student.GET("/iquireSuggestion", studentController.InquireSuggestion)
		student.GET("/deleteSuggestion", studentController.DeleteSuggestion)
		student.POST("/creatappeal", studentController.CreatAppeal)
	}
}
