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

type Suggestion struct {
	Advice string `json:"suggestion"`
	Name   string `json:"adminName"`
	Anon1  string `json:"anon"`
}

func PulishSuggest(c *gin.Context) {
	var data Suggestion
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
	adv := models.Advice{UserName: user.Name, Suggestion: data.Advice, Time: time.Now(), AdminId: teacher.AdminID, Anon: data.Anon1}
	error := database.DB.Create(&adv).Error
	if error != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
