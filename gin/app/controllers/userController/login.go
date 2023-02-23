package userController

import (
	"github.com/gin-gonic/gin"
	"log"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
)

type LoginData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
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
		return
	}
	err = sessionService.SetStudentSession(c, user)
	if err != nil {
		log.Println(err)
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonResponse(200, "ok", gin.H{
		"userid": user.ID,
		"admin":  user.Admin,
	}, c)
}
