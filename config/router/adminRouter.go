package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/adminController"
	"scoresystem/app/midwares"
)

func adminRouterInit(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	{
		admin.POST("/judge", adminController.Jugde, midwares.CheckLogin)
		admin.GET("/getunprocessedapplication", adminController.GetUnprocessedApplication, midwares.CheckLogin)
		admin.GET("/getreapprovedapplication", midwares.CheckLogin, adminController.GetReapprovedApplication)
		admin.POST("/putscore", midwares.CheckLogin, adminController.Putscore)
		admin.POST("/pulishReason", midwares.CheckLogin, adminController.PulishReason)
		admin.GET("/inquireReason", midwares.CheckLogin, adminController.InquireReason)
		admin.POST("/updateReason", midwares.CheckLogin, adminController.UpdateReason)
		admin.GET("/deleteReason", midwares.CheckLogin, adminController.DeleteReason)
		admin.GET("/inquireAdmin", adminController.InquireAdmin)
		admin.POST("/revokeapplication", midwares.CheckLogin, adminController.RevokeApplication)
	}
}
