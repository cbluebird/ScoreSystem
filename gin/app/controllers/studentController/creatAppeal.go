package studentController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"time"

	"github.com/gin-gonic/gin"
)

type AppealData struct {
	ApplicationID int    `json:"applicationid"`
	AppealReason  string `json:"reason"`
	Userid        int    `json:"userid"`
}

func CreatAppeal(c *gin.Context) {
	var appealdata AppealData
	var applicationdata models.Application
	err := c.ShouldBindJSON(&appealdata)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	_, err1 := userService.GetUserById(appealdata.Userid)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	appeal := &models.Appeal{
		ApplicationID: appealdata.ApplicationID,
		AppealReason:  appealdata.AppealReason,
		CreatTime:     time.Now(),
	}
	result := database.DB.Create(appeal).Error
	if result != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	result = database.DB.Where(&models.Application{ID: appeal.ApplicationID}).First(&applicationdata).Error
	if result != nil {
		utility.JsonResponse(405, "未找到相应的申请", nil, c)
		return
	}
	if applicationdata.Sta != 2 {
		utility.JsonResponse(404, "该申请状态未可申诉", nil, c)
		return
	}
	err2 := database.DB.Model(&models.Application{}).Where("id = ?", applicationdata.ID).Update("sta", "3").Error
	if err2 != nil {
		utility.JsonResponse(406, "修改申请状态错误", nil, c)
		return
	}
	utility.JsonResponse(200, "申诉发送成功", nil, c)
}
