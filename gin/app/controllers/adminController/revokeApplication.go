package adminController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/models/modules"
	"scoresystem/app/services/scoreService"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"scoresystem/config/database"

	"github.com/gin-gonic/gin"
)

type AppIDdata struct {
	ID     int `json:"applicationid"`
	Userid int `json:"userid"`
}

func RevokeApplication(c *gin.Context) {
	var applicationIDdata AppIDdata
	var applicationdata models.Application
	err := c.ShouldBindJSON(&applicationIDdata)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	user, err1 := userService.GetUserById(applicationIDdata.Userid)
	if err1 != nil || user.Admin == 0 {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	database.DB.Where(&models.Application{ID: applicationIDdata.ID}).First(&applicationdata)
	if applicationdata.Sta == 5 {
		utility.JsonResponse(404, "该申请未审批", nil, c)
		return
	}
	if applicationdata.Sta == 1 {
		if applicationdata.Module == 1 {
			var art modules.Art
			database.DB.Where(&modules.Art{ApplicationID: applicationIDdata.ID}).First(&art)
			res := database.DB.Delete(&art)
			if res.Error != nil {
				utility.JsonResponse(405, "对应数据删除错误", nil, c)
				return
			}
		}
		if applicationdata.Module == 2 {
			var gpa modules.GPA
			database.DB.Where(&modules.GPA{ApplicationID: applicationIDdata.ID}).First(&gpa)
			res := database.DB.Delete(&gpa)
			if res.Error != nil {
				utility.JsonResponse(405, "对应数据删除错误", nil, c)
				return
			}
		}
		if applicationdata.Module == 3 {
			var innovate modules.Innovate
			database.DB.Where(&modules.Innovate{ApplicationID: applicationIDdata.ID}).First(&innovate)
			res := database.DB.Delete(&innovate)
			if res.Error != nil {
				utility.JsonResponse(405, "对应数据删除错误", nil, c)
				return
			}
		}
		if applicationdata.Module == 4 {
			var labour modules.Labour
			database.DB.Where(&modules.Labour{ApplicationID: applicationIDdata.ID}).First(&labour)
			res := database.DB.Delete(&labour)
			if res.Error != nil {
				utility.JsonResponse(405, "对应数据删除错误", nil, c)
				return
			}
		}
		if applicationdata.Module == 5 {
			var moral modules.Moral
			database.DB.Where(&modules.Moral{ApplicationID: applicationIDdata.ID}).First(&moral)
			res := database.DB.Delete(&moral)
			if res.Error != nil {
				utility.JsonResponse(405, "对应数据删除错误", nil, c)
				return
			}
		}
		if applicationdata.Module == 6 {
			var pe modules.Pe
			database.DB.Where(&modules.Pe{ApplicationID: applicationIDdata.ID}).First(&pe)
			res := database.DB.Delete(&pe)
			if res.Error != nil {
				utility.JsonResponse(405, "对应数据删除错误", nil, c)
				return
			}
		}
	}
	err2 := database.DB.Model(&models.Application{}).Where("id = ?", applicationdata.ID).Update("sta", "5").Error
	if err2 != nil {
		utility.JsonResponse(406, "修改申请状态错误", nil, c)
		return
	}
	scoreService.UpdateScore(applicationdata.Userid, applicationdata.Age)
	utility.JsonResponse(200, "撤销成功", nil, c)
}
