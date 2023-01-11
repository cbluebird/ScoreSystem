package userController

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/apiExpection"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
)

type Pass struct {
	Password string `json:"password"`
}

func RePassword(c *gin.Context) {
	var data Pass
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	user, err1 := sessionService.GetUserSession(c)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	err = userService.UpdatePasswordByID(user, data.Password)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	sessionService.ClearUserSession(c)
	utility.JsonSuccessResponse(c, nil)
}
