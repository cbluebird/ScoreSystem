package user

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
)

type LoginData struct {
	Account  string
	Password string
}

func login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	user, err1 := userService.GetUserByAccount(data.Account)
	if err1 != nil {
		utility.JsonResponse(404, "用户不存在", nil, c)
		return
	}
	flag := userService.CheckUserByPassword(user, data.Password)
	if !flag {
		utility.JsonResponse(405, "密码错误", nil, c)
	}
	sessionService.SetStudentSession(c, user)
	utility.JsonSuccessResponse(c, nil)
}
