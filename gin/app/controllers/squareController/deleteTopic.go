package squareController

import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteTopic(c *gin.Context) {
	data := c.Query("squareId")
	userid, _ := strconv.Atoi(c.Query("userid"))
	user, err1 := userService.GetUserById(userid)
	if data == "" {
		utility.JsonResponseInternalServerError(c)
		return
	}
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	Int, _ := strconv.Atoi(data)
	topic := models.SquareTopic{}
	result1 := database.DB.Where(
		&models.SquareTopic{
			SquareId: Int,
		},
	).Find(&topic)
	if result1.Error != nil {
		utility.JsonResponse(406, "not find Topic", nil, c)
		return
	} //通过squareid找到对应topic
	if topic.UserId != user.ID {
		utility.JsonResponse(406, "can't delete the others topic", nil, c)
		return
	}
	result2 := database.DB.Delete(
		&models.SquareTopic{
			SquareId: Int,
			UserId:   user.ID,
		},
	)
	if result2.Error != nil {
		utility.JsonResponse(406, "not find topic", nil, c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
