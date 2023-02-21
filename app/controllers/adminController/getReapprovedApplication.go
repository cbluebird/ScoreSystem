package adminController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReapprovedApplication(c *gin.Context) {
	var application []models.Application
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if page <= 0 {
		page = 1
	}
	user, err1 := sessionService.GetUserSession(c)
	if err1 != nil || user.Admin == 0 {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	error := database.DB.Where(&models.Application{
		AdminId: user.ID,
		Sta:     3,
	}).Limit(pageSize).Offset((page - 1) * pageSize).Find(&application).Error
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
		"total":       total,
		"page":        page,
		"pagesize":    pageSize,
	}, c)
}
