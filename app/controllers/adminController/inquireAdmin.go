package adminController

import (
	"scoresystem/app/models"
	"scoresystem/app/utility"
	"scoresystem/config/database"

	"github.com/gin-gonic/gin"
)

type adminList struct {
	TeacherName string
	TeacherId   int
}

func InquireAdmin(c *gin.Context) {

	var sq []models.User
	error := database.DB.Where(&models.User{
		Admin: 1,
	}).Find(&sq).Error
	list := make([]adminList, len(sq))
	for i, value := range sq {
		list[i].TeacherName = value.Name
		list[i].TeacherId = value.ID
	}
	if error != nil {
		utility.JsonResponse(406, "not find admin", nil, c)
		return
	}
	utility.JsonResponse(200, "ok", gin.H{
		"teacherlist": list,
	}, c)
}
