package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/adminController"
)

func adminRouterInit(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	{
		admin.POST("/judge", adminController.Jugde)
		admin.GET("/getunprocessedapplication", adminController.GetUnprocessedApplication)
		admin.GET("/getreapprovedapplication", adminController.GetReapprovedApplication)
		admin.POST("/putscore", adminController.Putscore)
		admin.POST("/pulishReason", adminController.PulishReason)
		admin.GET("/inquireReason", adminController.InquireReason)
		admin.POST("/updateReason", adminController.UpdateReason)
		admin.GET("/deleteReason", adminController.DeleteReason)
		admin.GET("/inquireAdmin", adminController.InquireAdmin)
		admin.POST("/revokeapplication", adminController.RevokeApplication)
	}
}
