package userController

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/models"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
)

func CreateUser(c *gin.Context) {
	var data models.Reg
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	err = userService.CreateUser(data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
