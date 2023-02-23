package adminController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/utility"
	"scoresystem/config/database"

	"github.com/gin-gonic/gin"
)

func InquireReason(c *gin.Context) {

	user, err1 := sessionService.GetUserSession(c)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	var sq []models.Reason
	error := database.DB.Where(&models.Reason{
		AdminID: user.ID,
	}).Find(&sq).Error
	if error != nil {
		utility.JsonResponse(406, "not find Reason", nil, c)
		return
	}
	utility.JsonResponse(200, "ok", gin.H{
		"Reason": sq,
	}, c)
}
