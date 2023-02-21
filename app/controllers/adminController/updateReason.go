package adminController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/utility"
	"scoresystem/config/database"

	"github.com/gin-gonic/gin"
)

type Update struct {
	ID     int    `json:"id"`
	Reason string `json:"reason"`
}

func UpdateReason(c *gin.Context) {
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
	error := database.DB.Where(
		&models.Reason{
			Id:      data.ID,
			AdminID: user.ID,
		},
	).Updates(&models.Reason{Reason: data.Reason}).Error
	if error != nil {
		utility.JsonResponse(406, "can't update", nil, c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
