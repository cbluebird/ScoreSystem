package studentService

import (
	"scoresystem/app/models"
	"scoresystem/config/database"
	"time"
)

func Add(user *models.User, data models.ApplicationData) error {
	app := &models.Application{
		Age:         data.Age,
		Class:       data.Class,
		Module:      data.Module,
		CreateTime:  time.Now(),
		AdminId:     data.Admin,
		Description: data.Description,
		Sta:         5,
		Userid:      user.ID,
		Score:       data.Score,
	}
	result := database.DB.Create(app)
	return result.Error
}
