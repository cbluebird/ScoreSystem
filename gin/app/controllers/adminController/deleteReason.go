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

func DeleteReason(c *gin.Context) {
	data := c.Query("reasonId")
	if data == "" {
		utility.JsonResponseInternalServerError(c)
		return
	}
	user, err1 := sessionService.GetUserSession(c)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	Int, _ := strconv.Atoi(data)
	result := database.DB.Delete(
		&models.Reason{
			Id:      Int,
			AdminID: user.ID,
		},
	)
	if result.Error != nil {
		utility.JsonResponse(406, "can't delete Reason", nil, c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
