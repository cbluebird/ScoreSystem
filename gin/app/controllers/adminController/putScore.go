package adminController

import (
	"log"
	"scoresystem/app/apiExpection"
	"scoresystem/app/models"
	"scoresystem/app/models/modules"
	"scoresystem/app/services/scoreService"
	"scoresystem/app/services/userService"
	"scoresystem/app/utility"
	"scoresystem/config/database"
	"time"

	"github.com/gin-gonic/gin"
)

type Score struct {
	Score       float64 `json:"score"`
	Module      int     `json:"module"`
	StudentName string  `json:"studentname"`
	Class       int     `json:"class"`
	Type        int     `json:"type"`
	Age         int     `json:"age"`
	Userid      int     `json:"userid"`
}

func Putscore(c *gin.Context) {
	var data Score
	err := c.ShouldBindJSON(&data)
	log.Println(data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	user, err1 := userService.GetUserById(data.Userid)
	if err1 != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	student := models.User{}
	result0 := database.DB.Where(
		&models.User{
			Name: data.StudentName,
		},
	).Find(&student)
	if result0.Error != nil {
		utility.JsonResponse(406, "not find student", nil, c)
		return
	} //通过studentName找到对应studentid
	teacher := models.Admin{}
	result1 := database.DB.Where(
		&models.Admin{
			AdminID: user.ID,
		},
	).Find(&teacher)
	if result1.Error != nil {
		utility.JsonResponse(406, "not find admin", nil, c)
		return
	} //通过userid找到对应admin
	var flag int
	flag = 0
	if teacher.Art == 1 && data.Module == 1 {
		flag = 1
	}
	if teacher.Gpa == 1 && data.Module == 2 {
		flag = 1
	}
	if teacher.Innovate == 1 && data.Module == 3 {
		flag = 1
	}
	if teacher.Labour == 1 && data.Module == 4 {
		flag = 1
	}
	if teacher.Moral == 1 && data.Module == 5 {
		flag = 1
	}
	if teacher.PE == 1 && data.Module == 6 {
		flag = 1
	}
	if flag == 0 {
		utility.JsonResponse(406, "you don't have authority", nil, c)
		return
	}
	switch data.Module {
	case 1:
		{
			var art modules.Art
			if models.Authority[data.Module][data.Class] == 1 {
				return2 := database.DB.Where(&modules.Art{
					Userid: student.ID,
					Class:  data.Class,
				}).First(&art)
				if return2.RowsAffected == 1 {
					utility.JsonResponseInternalServerError(c)
					return
				}
			}
		}
	case 2:
		{
			var gpa modules.GPA
			if models.Authority[data.Module][data.Class] == 1 {
				return2 := database.DB.Where(&modules.GPA{
					Userid: student.ID,
					Class:  data.Class,
				}).First(&gpa)
				if return2.RowsAffected == 1 {
					utility.JsonResponseInternalServerError(c)
					return
				}

			}
		}
	case 3:
		{
			var innovate modules.Innovate
			if models.Authority[data.Module][data.Class] == 1 {
				return2 := database.DB.Where(&modules.Innovate{
					Userid: student.ID,
					Class:  data.Class,
				}).First(&innovate)
				if return2.RowsAffected == 1 {
					utility.JsonResponseInternalServerError(c)
					return
				}
			}
		}
	case 4:
		{
			var innovate modules.Labour
			if models.Authority[data.Module][data.Class] == 1 {
				return2 := database.DB.Where(&modules.Labour{
					Userid: student.ID,
					Class:  data.Class,
				}).First(&innovate)
				if return2.RowsAffected == 1 {
					utility.JsonResponseInternalServerError(c)
					return
				}
			}
		}
	case 5:
		{
			var moral modules.Moral
			if models.Authority[data.Module][data.Class] == 1 {
				return2 := database.DB.Where(&modules.Moral{
					Userid: student.ID,
					Class:  data.Class,
				}).First(&moral)
				if return2.RowsAffected == 1 {
					utility.JsonResponseInternalServerError(c)
					return
				}
			}
		}
	case 6:
		{
			var pe modules.Pe
			if models.Authority[data.Module][data.Class] == 1 {
				return2 := database.DB.Where(&modules.Pe{
					Userid: student.ID,
					Class:  data.Class,
				}).First(&pe)
				if return2.RowsAffected == 1 {
					utility.JsonResponseInternalServerError(c)
					return
				}
			}
		}
	}
	switch data.Module {
	case 1:
		{
			crt := modules.Art{Description: "模块成绩",Age: data.Age, Class: data.Class, CreateTime: time.Now(), Userid: student.ID, Score: float64(data.Score), ApplicationID: 0}
			error := database.DB.Create(&crt).Error
			if error != nil {
				utility.JsonResponseInternalServerError(c)
				return
			}
		}
	case 2:
		{
			crt := modules.GPA{Description: "模块成绩",Age: data.Age, Class: data.Class, CreateTime: time.Now(), Userid: student.ID, Score: float64(data.Score), ApplicationID: 0}
			error := database.DB.Create(&crt).Error
			if error != nil {
				utility.JsonResponseInternalServerError(c)
				return
			}
		}
	case 3:
		{
			crt := modules.Innovate{Description: "模块成绩",Age: data.Age, Class: data.Class, CreateTime: time.Now(), Userid: student.ID, Score: float64(data.Score), ApplicationID: 0}
			error := database.DB.Create(&crt).Error
			if error != nil {
				utility.JsonResponseInternalServerError(c)
				return
			}
		}
	case 4:
		{
			crt := modules.Labour{Description: "模块成绩",Age: data.Age, Class: data.Class, CreateTime: time.Now(), Userid: student.ID, Score: float64(data.Score), ApplicationID: 0}
			error := database.DB.Create(&crt).Error
			if error != nil {
				utility.JsonResponseInternalServerError(c)
				return
			}
		}
	case 5:
		{
			crt := modules.Moral{Description: "模块成绩",Age: data.Age, Class: data.Class, CreateTime: time.Now(), Userid: student.ID, Score: float64(data.Score), ApplicationID: 0}
			error := database.DB.Create(&crt).Error
			if error != nil {
				utility.JsonResponseInternalServerError(c)
				return
			}
		}
	case 6:
		{
			crt := modules.Pe{Description: "模块成绩",Age: data.Age, Class: data.Class, CreateTime: time.Now(), Userid: student.ID, Score: float64(data.Score), ApplicationID: 0}
			error := database.DB.Create(&crt).Error
			if error != nil {
				utility.JsonResponseInternalServerError(c)
				return
			}
		}
	}
	scoreService.UpdateScore(student.ID, data.Age)
	utility.JsonSuccessResponse(c, nil)
}
