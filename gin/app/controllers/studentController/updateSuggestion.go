package studentController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"time"

	"github.com/gin-gonic/gin"
)

type Update struct {
	ID     int    `json:"suggestionid"`
	Advice string `json:"suggestion"`
	Name   string `json:"adminName"`
	Anon1  string `json:"anon"`
}

func UpdateSuggestion(c *gin.Context) {
	var data Update
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	user, err1 := sessionService.GetUserSession(c)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	teacher := models.Admin{}
	result := database.DB.Where(
		&models.Admin{
			AdminName: data.Name,
		},
	).Find(&teacher)
	if result.Error != nil {
		utility.JsonResponse(406, "not find Admin", nil, c)
		return
	} //通过辅导员名字找到对应ID
	advice := models.Advice{}
	result1 := database.DB.Where(
		&models.Advice{
			Id: data.ID,
		},
	).Find(&advice)
	if result1.Error != nil {
		utility.JsonResponse(406, "not find advice", nil, c)
		return
	} //通过id找到对应advice
	if advice.UserName != user.Name {
		utility.JsonResponse(406, "can't update the others advice", nil, c)
		return
	}
	error := database.DB.Where(
		&models.Advice{
			Id:       data.ID,
			UserName: user.Name,
		},
	).Updates(&models.Advice{Time: time.Now(), Suggestion: data.Advice, AdminId: teacher.AdminID, UserName: user.Name, Anon: data.Anon1}).Error
	if error != nil {
		utility.JsonResponse(406, "can't update", nil, c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
