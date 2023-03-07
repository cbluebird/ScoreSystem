package ruleConrtoller

import (
	"github.com/gin-gonic/gin"
	"log"
	"scoresystem/app/models"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"
	"time"
)

type rule struct {
	Collage string `json:"collage"`
	Rule    string `json:"rule"`
}

func Insert(c *gin.Context) {
	var data rule
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	var r1 models.Rule
	result := database.DB.Where(&models.Rule{Collage: data.Collage}).First(&r1)
	if result.RowsAffected != 0 {
		utility.JsonResponse(400, "规则已存在", nil, c)
		return
	}
	r := &models.Rule{
		Rule:       data.Rule,
		CreateTime: time.Now(),
		Collage:    data.Collage,
	}
	database.DB.Create(&r)
	utility.JsonSuccessResponse(c, nil)
}

func Delete(c *gin.Context) {
	ruleId_ := c.Query("rule")
	ruleid, _ := strconv.Atoi(ruleId_)
	log.Println(ruleid)
	database.DB.Delete(&models.Rule{ID: ruleid})
	utility.JsonSuccessResponse(c, nil)
}

func Get(c *gin.Context) {
	collage := c.Query("collage")
	var r models.Rule
	database.DB.Where(&models.Rule{Collage: collage}).First(&r)
	utility.JsonSuccessResponse(c, r)
}
