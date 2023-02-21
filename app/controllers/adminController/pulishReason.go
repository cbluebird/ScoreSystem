package adminController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/utility"
	"scoresystem/config/database"

	"github.com/gin-gonic/gin"
)

type Reason struct {
	Reason string `json:"reason"`
}

func PulishReason(c *gin.Context) {
	var data Reason
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
	adv := models.Reason{Reason: data.Reason, AdminID: user.ID}
	error := database.DB.Create(&adv).Error
	if error != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
