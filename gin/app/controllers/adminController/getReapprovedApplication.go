package adminController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReApplication struct {
	ID          int `gorm:"primary_key;AUTO_INCREMENT"`
	Age         int
	Module      int
	Class       int
	CreateTime  time.Time
	Userid      int
	Score       float32
	Description string
	Sta         int
	AdminId     int
	Reason      string
	Appeal      string
}

func GetReapprovedApplication(c *gin.Context) {
	var application []models.Application
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	userid, _ := strconv.Atoi(c.Query("userid"))
	user, err1 := userService.GetUserById(userid)
	if page <= 0 {
		page = 1
	}
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
		"total":       len(application),
		"page":        page,
		"pagesize":    pageSize,
	}, c)
}
