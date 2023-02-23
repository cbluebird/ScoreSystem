package router

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/controllers/userController"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/createuser", userController.CreateUser)
		user.POST("/login", userController.Login)
		bind := user.Group("/bind")
		{
			bind.POST("/repassword", userController.RePassword)
		}
	}
}
