package squareController

//用户查询已提交的申报表
import (
	"scoresystem/app/models"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InquireTopic(c *gin.Context) {

	PageNum := c.Query("page_num")
	PageSize := c.Query("page_size")
	PN, _ := strconv.Atoi(PageNum)
	PS, _ := strconv.Atoi(PageSize)
	if PN <= 0 {
		PN = 1
	}
	var sq []models.SquareTopic

	error := database.DB.Where(&models.SquareTopic{}).Limit(PS).Offset((PN - 1) * PS).Order("Time desc").Find(&sq).Error
	if error != nil {
		utility.JsonResponse(406, "not find squaretopic", nil, c)
		return
	}
	var total int64
	if err := database.DB.Model(&models.SquareTopic{}).Count(&total).Error; err != nil {
		utility.JsonResponse(406, "not find squaretopic", nil, c)
		return
	}
	utility.JsonResponse(200, "ok", gin.H{
		"square":   sq,    //话题广场内容
		"total":    total, //总数
		"page":     PN,    //分页页数
		"pagesize": PS,    //分页大小
	}, c)
}
