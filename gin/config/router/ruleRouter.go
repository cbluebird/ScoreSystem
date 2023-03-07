package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/ruleConrtoller"
)

func ruleRouterInit(r *gin.RouterGroup) {
	rule := r.Group("/rule")
	rule.GET("/get", ruleConrtoller.Get)
	rule.POST("/add", ruleConrtoller.Insert)
	rule.GET("/delete",ruleConrtoller.Delete)
}
