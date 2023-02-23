package squareController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"time"

	"github.com/gin-gonic/gin"
)

type Topic struct {
	Word   string `json:"word"`
	Userid int    `json:"userid"`
}

func PulishTopic(c *gin.Context) {
	var data Topic
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	user, err1 := userService.GetUserById(data.Userid)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	square := models.SquareTopic{UserId: user.ID, Word: data.Word, Time: time.Now()}
	error := database.DB.Create(&square).Error
	if error != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
