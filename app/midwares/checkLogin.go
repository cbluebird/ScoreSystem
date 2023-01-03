package midwares

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/apiExpection"
	"scoresystem/app/services/sessionService"
)

func CheckLogin(c *gin.Context) {
	isLogin := sessionService.CheckUserSession(c)
	if !isLogin {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	c.Next()
}
