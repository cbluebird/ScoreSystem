package adminController

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/apiExpection"
	"scoresystem/app/services/adminService"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/utility"
)

type JudgeData struct {
	Status int `json:"status"`
	Id     int `json:"id"`
}

func Jugde(c *gin.Context) {
	var data JudgeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	user, err1 := sessionService.GetUserSession(c)
	if err1 != nil || user.Admin == 0 {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	flag := adminService.GetApp(data.Id, user)
	if !flag {
		utility.JsonResponse(404, "找不到对应的申请", nil, c)
	}
	err = adminService.Judge(data.Status, data.Id)
	utility.JsonSuccessResponse(c, nil)
}
