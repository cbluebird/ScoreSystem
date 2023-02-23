package studentController

//用户查询已提交的申报表
import (
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/services/sessionService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InquireSuggestion(c *gin.Context) {
	_, err1 := sessionService.GetUserSession(c)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	PageNum := c.Query("page_num")
	PageSize := c.Query("page_size")
	PN, _ := strconv.Atoi(PageNum)
	PS, _ := strconv.Atoi(PageSize)
	if PN <= 0 {
		PN = 1
	}
	var sq []models.Advice

	error := database.DB.Where(&models.Advice{}).Limit(PS).Offset((PN - 1) * PS).Order("Time desc").Find(&sq).Error
	if error != nil {
		utility.JsonResponse(406, "not find advice", nil, c)
		return
	}
	var total int64
	if err := database.DB.Model(&models.Advice{}).Count(&total).Error; err != nil {
		utility.JsonResponse(406, "not find advice", nil, c)
		return
	}

	utility.JsonResponse(200, "ok", gin.H{
		"advice":   sq,    //建议内容
		"total":    total, //总数
		"page":     PN,    //分页页数
		"pagesize": PS,    //分页大小
	}, c)
}
