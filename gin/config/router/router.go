package router

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	const pre = "/api"

	api := r.Group(pre)
	{
		userRouterInit(api)
		studentRouterInit(api)
		adminRouterInit(api)
		squareRouterInit(api)
		ruleRouterInit(api)
	}
}
