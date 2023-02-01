package adminService

import (
	"scoresystem/app/models"
	"scoresystem/app/models/modules"
	"scoresystem/config/database"
	"time"
)

func Judge(status, id int) error {
	app := models.Application{}
	result := database.DB.Where(
		&models.User{
			ID: id,
		},
	).First(&app)
	if result.Error != nil {
		return result.Error
	}
	app.Status = status
	err := database.DB.Model(models.Application{}).Where(
		models.Application{
			ID: id,
		}).Updates(app).Error
	if err != nil {
		return err
	}
	if status == 2 || status == 4 {
		return nil
	}
	switch app.Module {
	case 1:
		{
			data := &modules.Art{
				Age:         app.Age,
				Class:       app.Class,
				CreateTime:  time.Now(),
				Userid:      app.Userid,
				Score:       app.Score,
				Description: app.Description,
				Type:        0,
			}
			result = database.DB.Create(data)
			err = result.Error
		}
	case 2:
		{
			data := &modules.GPA{
				Age:         app.Age,
				Class:       app.Class,
				CreateTime:  time.Now(),
				Userid:      app.Userid,
				Score:       app.Score,
				Description: app.Description,
				Type:        0,
			}
			result = database.DB.Create(data)
			err = result.Error
		}
	case 3:
		{
			data := &modules.Innovate{
				Age:         app.Age,
				Class:       app.Class,
				CreateTime:  time.Now(),
				Userid:      app.Userid,
				Score:       app.Score,
				Description: app.Description,
				Type:        0,
			}
			result = database.DB.Create(data)
			err = result.Error
		}
	case 4:
		{
			data := &modules.Labour{
				Age:         app.Age,
				Class:       app.Class,
				CreateTime:  time.Now(),
				Userid:      app.Userid,
				Score:       app.Score,
				Description: app.Description,
				Type:        0,
			}
			result = database.DB.Create(data)
			err = result.Error
		}
	case 5:
		{
			data := &modules.Moral{
				Age:         app.Age,
				Class:       app.Class,
				CreateTime:  time.Now(),
				Userid:      app.Userid,
				Score:       app.Score,
				Description: app.Description,
				Type:        0,
			}
			result = database.DB.Create(data)
			err = result.Error
		}
	case 6:
		{
			data := &modules.Pe{
				Age:         app.Age,
				Class:       app.Class,
				CreateTime:  time.Now(),
				Userid:      app.Userid,
				Score:       app.Score,
				Description: app.Description,
				Type:        0,
			}
			result = database.DB.Create(data)
			err = result.Error
		}
	}
	if err != nil {
		return result.Error
	} else {
		return nil
	}
}

func GetApp(id int, user *models.User) bool {
	app := models.Application{}
	result := database.DB.Where(
		&models.User{
			ID: id,
		},
	).First(&app)
	if result.Error != nil || app.AdminId != user.ID {
		return false
	}
	return true
}
