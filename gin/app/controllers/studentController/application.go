package studentController

import (
	"github.com/gin-gonic/gin"
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/studentService"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
)

func Add(c *gin.Context) {
	var data models.ApplicationData
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
	err = studentService.Add(user, data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
