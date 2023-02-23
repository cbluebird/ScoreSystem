package studentController

import (
	"github.com/gin-gonic/gin"
	"log"
	"scoresystem/app/apiExpection"
	"scoresystem/app/services/scoreService"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"strconv"
)

type responseData struct {
	moudle int
	class  int
	score  float32
}

func GetAll(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Query("userid"))
	user, err := userService.GetUserById(userid)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	year_ := c.Query("year")
	year, _ := strconv.Atoi(year_)
	list, _ := scoreService.GetAllScore(user.ID, year)
	log.Println(list)
	utility.JsonResponse(200, "ok", gin.H{
		"Score": list.Grade,
	}, c)
}

func GetDetail(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Query("userid"))
	user, err := userService.GetUserById(userid)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	year_ := c.Query("year")
	year, _ := strconv.Atoi(year_)
	list, _ := scoreService.GetDetail(user.ID, year)
	utility.JsonResponse(200, "ok", gin.H{
		"detaillist": list,
	}, c)
}
