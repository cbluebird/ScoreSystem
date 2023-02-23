package adminController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUnprocessedApplication(c *gin.Context) {
	var application []models.Application
	pageNum, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageNum <= 0 {
		pageNum = 1
	}
	userid, _ := strconv.Atoi(c.Query("userid"))
	user, err := userService.GetUserById(userid)
	if err != nil || user.Admin == 0 {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	error := database.DB.Where(&models.Application{AdminId: user.ID, Sta: 5}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&application).Error
	if error != nil {
		utility.JsonResponse(406, "not find application", nil, c)
		return
	}
	var total int64
	if err := database.DB.Model(&models.Application{}).Count(&total).Error; err != nil {
		utility.JsonResponse(406, "not find application", nil, c)
		return
	}
	utility.JsonResponse(200, "ok", gin.H{
		"application": application,
	}, c)
}
