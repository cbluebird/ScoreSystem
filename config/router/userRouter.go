package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/userController"
	"scoresystem/app/midwares"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/createuser", userController.CreateUser)
		user.POST("/login", userController.Login)
		bind := user.Group("/bind", midwares.CheckLogin)
		{
			bind.POST("/repassword", userController.RePassword)
		}
	}
}
