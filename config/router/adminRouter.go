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
	}
}
