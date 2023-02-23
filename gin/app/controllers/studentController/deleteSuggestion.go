package studentController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteSuggestion(c *gin.Context) {
	data := c.Query("suggestionId")
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
	advice := models.Advice{}
	result1 := database.DB.Where(
		&models.Advice{
			Id: Int,
		},
	).Find(&advice)
	if result1.Error != nil {
		utility.JsonResponse(406, "not find advice", nil, c)
		return
	} //通过suggestionid找到对应advice
	if advice.UserName != user.Name {
		utility.JsonResponse(406, "can't delete the others advice", nil, c)
		return
	}
	result := database.DB.Delete(
		&models.Advice{
			Id:       Int,
			UserName: user.Name,
		},
	)
	if result.Error != nil {
		utility.JsonResponse(406, "not find Advice", nil, c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
