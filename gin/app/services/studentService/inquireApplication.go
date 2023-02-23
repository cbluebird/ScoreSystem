package studentService

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InquireApplication(c *gin.Context) {
	var application []models.Application
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if page <= 0 {
		page = 1
	}
	userid, _ := strconv.Atoi(c.Query("userid"))
	user, err1 := userService.GetUserById(userid)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	err2 := database.DB.Where(&models.Application{
		Userid: user.ID,
	}).Limit(pageSize).Offset((page - 1) * pageSize).Find(&application).Error
	if err2 != nil {
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
		"total":       len(application),
		"page":        page,
		"pagesize":    pageSize,
	}, c)
}
